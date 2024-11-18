package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

type NetLogger struct {
	Network   *NetStatus
	LogBuffer []*http.Response
}

type LocalLogger struct {
	FilePtr   *os.File
}

type Logger interface {
	Add(res *http.Response)
	Commit()
	Dispose()
}

////////////////////////////////////////////////////////////// NetLogger implemantation
func initNetLogger(netStatus *NetStatus) *NetLogger {
	logger := new(NetLogger)
	logger.Network = netStatus
	return logger
}

func (logger *NetLogger) Add(res *http.Response) {
	logger.LogBuffer = append(logger.LogBuffer, res)
}

func (logger *NetLogger) Commit() {
	if !logger.Network.IsHelper {
		return
	}

	sendLog(logger.Network.Parent, logger.LogBuffer)
	logger.LogBuffer = nil
}

func (logger *NetLogger) Dispose() {
	// Do nothing, we can't close the connection here
	// it needs to be used elsewhere
}


////////////////////////////////////////////////////////////// LocalLogger implemantation
func initLocalLogger(flags *Flags) *LocalLogger {
	filePtr, err := os.OpenFile(flags.OutFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("ERROR: Can't create/open output file")
		return nil
	}

	logger := new(LocalLogger)
	logger.FilePtr = filePtr

	return logger
}

func (logger *LocalLogger) Add(res *http.Response) {
	var log strings.Builder

	//log.WriteString(fmt.Sprintf("--- [Request with param %s set to %s]", paramName, paramValue))
	log.WriteString("---------------------- [START OF HTTP RESPONSE] ----------------------\n")

	raw, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Println("ERROR: Couldn't log response, skipping it")
		return
	}
	log.WriteString(string(raw))

	log.WriteString("\n----------------------- [END OF HTTP RESPONSE] -----------------------\n\n")

	_, err = logger.FilePtr.WriteString(log.String())
	if err != nil {
		fmt.Println("ERROR: Couldn't log response, error wrtting to file, skipping it")
	}
}

func (logger *LocalLogger) Commit() {
	// Completely unnecesary but slightly better than having an empty implementation
	logger.FilePtr.Sync()
}

func (logger *LocalLogger) Dispose() {
	logger.FilePtr.Close()
}

//// Not in interface
func (logger *LocalLogger) BeginLog() {
	logger.FilePtr.WriteString("============================ [START OF SESSION] ============================\n")
}

func (logger *LocalLogger) EndLog() {
	logger.FilePtr.WriteString("============================= [END OF SESSION] =============================\n")
}
