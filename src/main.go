package main

import (
	"fmt"
	"net/http"
	"os"
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

		runDict(config, flags, paramNames[0], dict)
		return
	}

	fmt.Println("ERROR: No valid param type found, aborting")
}

func runDict(config *Config, flags *Flags, paramName string, dict []string) {
	dictCount := len(dict)
	batches := dictCount / flags.BatchSize
	if batches < 1 {
		batches = 1
	}

	for i := 0; i < batches; i++ {
		var waitGroup sync.WaitGroup

		batchStart := i * flags.BatchSize
		batchEnd := (i + 1) * flags.BatchSize
		if batchEnd > dictCount {
			batchEnd = dictCount
		}

		for _, elem := range dict[batchStart:batchEnd] {
			waitGroup.Add(1)
			go func(val string) {
				fmt.Printf("[%s]: ", val)
				makeRequest(config.Request, paramName, val)
				waitGroup.Done()
			}(elem)
		}

		waitGroup.Wait()
	}

}

func runRange(config *Config, flags *Flags, netStat *NetStatus, paramName string, from int, to int) {
	count := to - from
	batches := count / flags.BatchSize
	if batches < 1 {
		batches = 1
	}

	stop := config.Criteria.Type == "STOP"

	iter := from
	for batch := 1; batch <= batches; batch++ {
		responseCh := make(chan *http.Response, min(flags.BatchSize, count))	

		for ; iter <= to; iter++ {
			go iterFunc(config, paramName, strconv.Itoa(iter), responseCh)
			
			if iter%flags.BatchSize == 0 {
				break
			}
		}

		// TODO: Finish this
		for i := 0; i < cap(responseCh); i++ {
			res := <- responseCh
	
			if res != nil {
				config.Logger.Add(res)
	
				if stop {
					config.Logger.Commit()
					fmt.Println("STOP STOP STOP")
					return
				}
			}
		}

		config.Logger.Commit()
	}
}

func iterFunc(config *Config, paramName string, value string, channel chan *http.Response) {
	res := makeRequest(config.Request, paramName, value)
	ok := checkCriteria(res, &config.Criteria.Response)

	if ok && config.Criteria.Type != "" {
		channel <- res
	} else {
		channel <- nil
	}
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