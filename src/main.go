package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"qdbf/distributed"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var config *Config
	var err error
	var localLogger *LocalLogger

	netStat := new(NetStatus)

	flags, helper := handleFlags()

	if helper {
		netStat.IsHelper = true
		config, netStat.Parent = helperMode(flags)
		config.Logger = initNetLogger(netStat)

	} else {
		config, err = loadConfig(flags.ConfigFile)
		if err != nil {
			fmt.Println("ERROR: Couln't load config file!, please check the format")
			return
		}

		if config.Helpers != nil {
			netStat.Helpers = connectToHelpers(config.Helpers, flags.Port)
		}

		localLogger = initLocalLogger(flags)
		localLogger.BeginLog()
		config.Logger = localLogger
	}

	if config == nil {
		fmt.Println("ERROR: Invalid config, aborting")
		return
	}

	paramNames := loadParams(config)
	start(config, flags, paramNames, netStat)

	if !helper {
		localLogger.EndLog()
	}
}

func start(config *Config, flags *Flags, paramNames []string, netStat *NetStatus) {
	firstParam := config.Params[paramNames[0]]

	if firstParam.Type == "RANGE" {
		fmt.Printf("Found RANGE [%d, %d] param: %s\n", firstParam.From, firstParam.To, paramNames[0])

		if !netStat.IsHelper && netStat.Helpers != nil && len(netStat.Helpers) > 0 {
			splitWorkRange(netStat, config, flags, &firstParam, paramNames[0])
		}

		runRange(config, flags, netStat, paramNames[0], firstParam.From, firstParam.To)
		return
	}

	if firstParam.Type == "FILE" {
		data, err := os.ReadFile(firstParam.File)
		if err != nil {
			fmt.Println("ERROR: Could not read file, aborting")
			return
		}

		lines := strings.Split(string(data), "\n")

		firstParam.Dict = lines
		firstParam.Type = "DICT"
	}

	if firstParam.Type == "DICT" {
		dict := firstParam.Dict
		fmt.Printf("Found DICT [%d] param: %s\n", len(dict), paramNames[0])

		if !netStat.IsHelper && netStat.Helpers != nil && len(netStat.Helpers) > 0 {
			dict = splitWorkDict(netStat, config, flags, dict, paramNames[0])
		}

		runDict(config, flags, netStat, paramNames[0], dict)
		return
	}

	fmt.Println("ERROR: No valid param type found, aborting")
}

func runDict(config *Config, flags *Flags, netStat *NetStatus, paramName string, dict []string) {
	dictCount := len(dict)
	batches := dictCount / flags.BatchSize
	if batches < 1 {
		batches = 1
	}

	stop := config.Criteria.Type == "STOP"
	ctx, cancel := context.WithCancel(context.Background())
	responseCh := new(ResponseChannel)

	var wg sync.WaitGroup
	wg.Add(1)
	go handleNetwork(config, netStat, responseCh, ctx, cancel, &wg)

	for i := 0; i < batches; i++ {
		ch := make(chan *http.Response, min(flags.BatchSize, dictCount))
		responseCh.Assign(ch)

		batchStart := i * flags.BatchSize
		batchEnd := (i + 1) * flags.BatchSize
		if batchEnd > dictCount {
			batchEnd = dictCount
		}

		for _, elem := range dict[batchStart:batchEnd] {
			go iterFunc(config, paramName, elem, responseCh, ctx)
		}

		if handleResponses(responseCh, config, stop, ctx, cancel, netStat) {
			return
		}

		config.Logger.Commit()
		responseCh.Close()
	}
	cancel()
}

func runRange(config *Config, flags *Flags, netStat *NetStatus, paramName string, from int, to int) {
	count := to - from
	batches := count / flags.BatchSize
	if batches < 1 {
		batches = 1
	}

	stop := config.Criteria.Type == "STOP"
	ctx, cancel := context.WithCancel(context.Background())
	responseCh := new(ResponseChannel)

	var wg sync.WaitGroup
	wg.Add(1)
	go handleNetwork(config, netStat, responseCh, ctx, cancel, &wg)

	iter := from
	for batch := 1; batch <= batches; batch++ {
		ch := make(chan *http.Response, min(flags.BatchSize, count))
		responseCh.Assign(ch)

		for ; iter <= to; iter++ {
			go iterFunc(config, paramName, strconv.Itoa(iter), responseCh, ctx)

			if iter%flags.BatchSize == 0 {
				break
			}
		}

		if handleResponses(responseCh, config, stop, ctx, cancel, netStat) {
			return
		}

		config.Logger.Commit()
		responseCh.Close()
	}
	cancel()
	wg.Wait()
}

func handleResponses(responseCh *ResponseChannel, config *Config, stop bool, ctx context.Context, cancel context.CancelFunc, netStat *NetStatus) bool {
	for i := 0; i < cap(responseCh.Channel); i++ {
		res, ok := <- responseCh.Channel
		if !ok {
			return true
		}

		select {
		case <-ctx.Done():
			return true
		default:
		}

		if res != nil {
			config.Logger.Add(res)

			if stop {
				cancel()
				config.Logger.Commit()
				responseCh.Close()

				fmt.Println("STOP: Criteria met, stopping...")
				netStat.NetStop()
				return true
			}
		}
	}

	return false
}

func handleNetwork(config *Config, netStat *NetStatus, responseCh *ResponseChannel, ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	// No helpers active, no network communication
	if !netStat.IsHelper && len(netStat.Helpers) == 0 {
		return
	}

	stop := false
	var messages []*distributed.SyncMessage

	for !stop {
		select {
		case <-ctx.Done():
			return
		default:	
		}

		if netStat.IsHelper {
			stop = netStat.CheckParentStop()
		} else {
			messages, stop = netStat.CheckHelpers()
			if len(messages) > 0 {
				handleRemoteLogs(messages, config)
			}
		}
	}

	cancel()
	responseCh.Close()

	fmt.Println("STOP: Recieved remote STOP signal!")

	// If main instance, send message to all helpers
	if !netStat.IsHelper {
		netStat.NetStop()
	}
}

func handleRemoteLogs(messages []*distributed.SyncMessage, config *Config) {
	for _, msg := range messages {
		if msg.GetAction() == "LOG" {
			for _, res := range msg.GetResponseLog().GetReponses() {
				httpRes := res2httpRes(res)
				config.Logger.Add(httpRes)
			}
		}
	}
}

func iterFunc(config *Config, paramName string, value string, resChannel *ResponseChannel, ctx context.Context) {
	res := makeRequest(config.Request, paramName, value, ctx)
	if res == nil {
		return
	}

	select {
	case <-ctx.Done():
		return
	default:
	}

	ok := checkCriteria(res, &config.Criteria.Response)

	if ok {
		resChannel.Channel <- res
	} else {
		resChannel.Channel <- nil
	}

	fmt.Printf("[%s: %s] %s\n", paramName, value, res.Status)
}

func splitWorkRange(netStat *NetStatus, config *Config, flags *Flags, param *Param, paramName string) {
	helperCount := len(netStat.Helpers) + 1
	paramRange := param.To - param.From
	step := paramRange / helperCount
	extra := paramRange % helperCount

	startFrom := param.From + step + extra
	param.To = startFrom
	for _, helper := range netStat.Helpers {
		helperConfig := *config

		helperParam := helperConfig.Params[paramName]
		helperParam.From = startFrom
		helperParam.To = startFrom + step
		helperConfig.Params[paramName] = helperParam

		sendStart(helper, &helperConfig, flags)

		startFrom += step
	}

	fmt.Println("Work split between", helperCount-1, "workers and this instance")
	fmt.Printf("New range [%d, %d] for RANGE param %s on this instance\n", param.From, param.To, paramName)
}

func splitWorkDict(netStat *NetStatus, config *Config, flags *Flags, dict []string, paramName string) []string {
	helperCount := len(netStat.Helpers) + 1
	dictCount := len(dict)
	partitionSize := dictCount / helperCount
	extra := dictCount % helperCount

	localDict := dict[0 : partitionSize+extra]
	for i, helper := range netStat.Helpers {
		helperConfig := *config

		helperParam := helperConfig.Params[paramName]
		from := partitionSize*(i+1) + extra
		helperParam.Dict = dict[from : from+partitionSize]
		helperConfig.Params[paramName] = helperParam

		sendStart(helper, &helperConfig, flags)
	}

	fmt.Println("Work split between", helperCount-1, "workers and this instance")
	fmt.Printf("New size [%d] for DICT param %s on this instance\n", partitionSize+extra, paramName)

	return localDict
}
