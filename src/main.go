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
			startHelpers(netStat.Helpers, config, flags)
		}	
	}

	if config == nil {
		fmt.Println("ERROR: Invalid config, aborting")
		return
	}

	paramNames := loadParams(config)
	fmt.Println("ADFGH:", paramNames[0])
	//start(config, flags, paramNames, netStat)
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
		runRange(config, flags, paramNames[0], firstParam.From, firstParam.To)
		return
	}
	if firstParam.Type == "DICT" {
		return
	}
}


func runRange(config *Config, flags *Flags, name string, from int, to int) {
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
				makeRequest(config.Request, name, strconv.Itoa(i))
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