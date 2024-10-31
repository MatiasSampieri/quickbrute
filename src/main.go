package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var config *Config
	var err error

	var netStat NetStatus

	flags, helper := handleFlags()

	if helper {
		netStat.isHelper = true
		config, netStat.Parent = helperMode(flags)

	} else {
		config, err = loadConfig(flags.ConfigFile)
		if err != nil {
			fmt.Println("ERROR: Couln't load config file!, please check the format")
			return
		}

		if config.Helpers != nil {
			netStat.Helpers = connectToHelpers(config.Helpers, flags.Port)
		}	
	}

	if config == nil {
		fmt.Println("ERROR: Invalid config, aborting")
		return
	}

	paramNames := loadParams(config)
	start(config, flags, paramNames, netStat)
}


func helperMode(flags *Flags) (*Config, net.Conn) {
	conn := waitForMainInstance(flags.Port)
	if conn == nil {
		return nil, nil
	}

	config := waitForRemoteStart(conn, flags)
	if config == nil {
		fmt.Println("ERROR: Couln't load remote config! aborting")
		return nil, nil
	}

	return config, conn
}


func start(config *Config, flags *Flags, paramNames []string, netStat NetStatus) {
	firstParam := config.Params[paramNames[0]]
	if firstParam.Type == "RANGE" {
		fmt.Printf("Found RANGE [%d, %d] param: %s\n", firstParam.From, firstParam.To, paramNames[0])

		if !netStat.isHelper && netStat.Helpers != nil && len(netStat.Helpers) > 0 {
			splitWorkRange(netStat, config, flags, &firstParam, paramNames[0])
		}

		runRange(config, flags, paramNames[0], firstParam.From, firstParam.To)
		return
	}
	if firstParam.Type == "DICT" {
		dict := firstParam.Dict
		fmt.Printf("Found DICT [%d] param: %s\n", len(dict), paramNames[0])

		if !netStat.isHelper && netStat.Helpers != nil && len(netStat.Helpers) > 0 {
			dict = splitWorkDict(netStat, config, flags, dict, paramNames[0])
		}

		runDict(config, flags, paramNames[0], dict)
		return
	}
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
		batchEnd := (i+1) * flags.BatchSize
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


func runRange(config *Config, flags *Flags, paramName string, from int, to int) {
	batches := (to-from) / flags.BatchSize
	if batches < 1 {
		batches = 1
	}

	iter := from
	for batch := 1; batch <= batches; batch++ {
		var waitGroup sync.WaitGroup

		for ; iter <= to; iter++ {
			waitGroup.Add(1)
			go func(i int) {
				makeRequest(config.Request, paramName, strconv.Itoa(i))
				waitGroup.Done()
			}(iter)

			if iter % flags.BatchSize == 0 {
				break
			}
		}

		waitGroup.Wait()
	}
}


func makeRequest(request Request, paramName string, paramValue string) {
	URL := strings.ReplaceAll(request.URL, fmt.Sprintf("$%s$", paramName), paramValue)
	body := strings.ReplaceAll(request.Body, fmt.Sprintf("$%s$", paramName), paramValue)

	preparedUrl, err := url.Parse(URL)
	if err != nil {
		return
	}

	// Parse params
	params := preparedUrl.Query()
	for name, value := range request.Params {
		val := strings.ReplaceAll(value, fmt.Sprintf("$%s$", paramName), paramValue)
		params.Add(name, val)
	}
	preparedUrl.RawQuery = params.Encode()

	// Create request
	req, err := http.NewRequest(request.Method, preparedUrl.String(), strings.NewReader(body))
	if err != nil {
		return
	}

	// Parse headers
	for name, value := range request.Headers {
		val := strings.ReplaceAll(value, fmt.Sprintf("$%s$", paramName), paramValue)
		req.Header.Set(name, val)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	fmt.Println(res.Status)
}


func splitWorkRange(netStat NetStatus, config *Config, flags *Flags, param *Param, paramName string) {
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


func splitWorkDict(netStat NetStatus, config *Config, flags *Flags, dict []string, paramName string) []string {
	helperCount := len(netStat.Helpers) + 1
	dictCount := len(dict)
	partitionSize := dictCount / helperCount
	extra := dictCount % helperCount

	localDict := dict[0:partitionSize+extra]
	for i, helper := range netStat.Helpers {
		helperConfig := *config

		helperParam := helperConfig.Params[paramName]
		from := partitionSize * (i+1) + extra
		helperParam.Dict = dict[from:from+partitionSize]
		helperConfig.Params[paramName] = helperParam

		sendStart(helper, &helperConfig, flags)
	}

	fmt.Println("Work split between", helperCount-1, "workers and this instance")
	fmt.Printf("New size [%d] for DICT param %s on this instance\n", partitionSize+extra, paramName)

	return localDict
}