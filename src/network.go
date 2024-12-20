package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"qdbf/distributed"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
)

func checkSync(conn net.Conn, timeout time.Duration) *distributed.SyncMessage {
	if timeout == time.Duration(0) {
		conn.SetReadDeadline(time.Time{})
	} else {
		conn.SetReadDeadline(time.Now().Add(time.Millisecond * timeout))
	}

	data := make([]byte, 4096)
	count, err := conn.Read(data)
	if err != nil {
		return nil
	}

	msg := distributed.SyncMessage{}
	err = proto.Unmarshal(data[:count], &msg)
	if err != nil {
		fmt.Println("Bad request!")
		return nil
	}

	return &msg
}

func sendSync(conn net.Conn, message *distributed.SyncMessage) error {
	resData, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	_, err = conn.Write(resData)
	return err
}

func sendAction(conn net.Conn, action string) error {
	return sendSync(conn, &distributed.SyncMessage{Action: action})
}

func waitForMainInstance(port int) net.Conn {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("ERROR: Could not start listening on port:", port)
		return nil
	}

	fmt.Println("Waiting for main instance to communicate...")

	var conn net.Conn
	for {
		conn, err = listener.Accept()
		if err != nil {
			fmt.Println("ERROR: Could not establish connection!")
			continue
		}

		break
	}

	fmt.Println("Main instance conneted! Wating for HELLO from", conn.RemoteAddr().String())

	msg := checkSync(conn, 5000)

	if msg.GetAction() == "HELLO" {
		fmt.Println("Recieved HELLO from main instance!")
		sendAction(conn, "ACK")
		fmt.Println("Main instance conneted succesfully!")
	} else {
		fmt.Println("ERROR: HELLO not recieved!", msg.GetAction())
		return nil
	}

	return conn
}

func waitForRemoteStart(conn net.Conn, flags *Flags) *Config {
	fmt.Println("Witing for main instance to send START signal...")

	var msg *distributed.SyncMessage
	for msg = checkSync(conn, 0); msg.GetAction() != "START"; msg = checkSync(conn, 0) {
	}

	// Parse remote config
	rc := msg.GetConfig()

	request := Request{
		Method:  rc.GetRequest().GetMethod(),
		URL:     rc.GetRequest().GetUrl(),
		Body:    rc.GetRequest().GetBody(),
		Headers: rc.GetRequest().GetHeaders(),
		Params:  rc.GetRequest().GetParams(),
	}

	criteria := Criteria{
		Type: rc.GetCriteria().GetType(),
		Response: Response{
			Status:  int(rc.GetCriteria().GetResponse().GetStatus()),
			Body:    rc.GetCriteria().GetResponse().GetBody(),
			Headers: rc.GetCriteria().GetResponse().GetHeaders(),
		},
	}

	params := make(map[string]Param, len(rc.GetParams()))
	for name, param := range rc.GetParams() {
		params[name] = Param{
			Type: param.GetType(),
			From: int(param.GetFrom()),
			To:   int(param.GetTo()),
			Dict: param.GetDict(),
		}
	}

	config := Config{
		Request:  request,
		Params:   params,
		Criteria: criteria,
	}

	//flags.BatchSize = int(msg.GetFlags().GetBatchSize())

	fmt.Println("Recieved config from main instance!")
	sendAction(conn, "ACK")

	return &config
}

func sendStart(conn net.Conn, config *Config, flags *Flags) bool {
	fmt.Println("Sending START signal to", conn.RemoteAddr().String())

	request := distributed.Request{
		Method:  config.Request.Method,
		Url:     config.Request.URL,
		Body:    config.Request.Body,
		Headers: config.Request.Headers,
		Params:  config.Request.Params,
	}

	criteria := distributed.Criteria{
		Type: config.Criteria.Type,
		Response: &distributed.Response{
			Status:  int32(config.Criteria.Response.Status),
			Body:    config.Criteria.Response.Body,
			Headers: config.Criteria.Response.Headers,
		},
	}

	params := make(map[string]*distributed.Params, len(config.Params))
	for name, param := range config.Params {
		params[name] = &distributed.Params{
			Type: param.Type,
			From: int32(param.From),
			To:   int32(param.To),
			Dict: param.Dict,
		}
	}

	helperConfig := distributed.Config{
		Request:  &request,
		Params:   params,
		Criteria: &criteria,
	}

	helperFlags := distributed.Flags{
		BatchSize: int32(flags.BatchSize),
	}

	startMessage := distributed.SyncMessage{
		Action: "START",
		Config: &helperConfig,
		Flags:  &helperFlags,
	}

	err := sendSync(conn, &startMessage)
	if err != nil {
		fmt.Println("Failed to send START")
		return false
	}

	fmt.Println("START sent, waiting for response...")

	msg := checkSync(conn, 5000)
	if msg == nil || msg.GetAction() != "ACK" {
		fmt.Println("Helper did not respond to START")
		return false
	}

	fmt.Println("Helper active!")

	return true
}

func askForHelp(addr string, port int) net.Conn {
	fmt.Println("Attempting connection with helper [timeout 2s]:", addr)

	dialer := net.Dialer{Timeout: time.Second * 2}
	conn, err := dialer.Dial("tcp", addr+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("ERROR: Could not connect to helper, skipping it")
		return nil
	}

	fmt.Println("Connected to helper, sending HELLO")
	sendAction(conn, "HELLO")

	res := checkSync(conn, 5000)
	if res.GetAction() != "ACK" {
		fmt.Println("ERROR: Did not recieved ACK from helper, aborting")
		conn.Close()
		return nil
	}

	fmt.Println("Successfully connected to helper!")

	return conn
}

func connectToHelpers(helpers []string, port int) []net.Conn {
	var connections []net.Conn

	for _, addr := range helpers {
		conn := askForHelp(addr, port)
		if conn != nil {
			connections = append(connections, conn)
		}
	}

	return connections
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

func sendLog(conn net.Conn, responses []*http.Response) {
	count := int32(len(responses))
	convertedResponses := make([]*distributed.Response, count)

	for i, res := range responses {
		body, _ := io.ReadAll(res.Body)

		convertedResponses[i] = &distributed.Response{
			Status:        int32(res.StatusCode),
			StatusTxt:     res.Status,
			ProtoVer:      res.Proto,
			Body:          string(body),
			Headers:       make(map[string]string),
			ContentLength: res.ContentLength,
		}

		for name, value := range res.Header {
			convertedResponses[i].Headers[name] = strings.Join(value, ", ")
		}
	}

	logMessage := distributed.SyncMessage{
		Action: "LOG",
		ResponseLog: &distributed.Log{
			Count:    count,
			Reponses: convertedResponses,
		},
	}

	sendSync(conn, &logMessage)
}

func res2httpRes(res *distributed.Response) *http.Response {
	httpRes := new(http.Response)
	httpRes.Status = res.GetStatusTxt()
	httpRes.StatusCode = int(res.GetStatus())
	httpRes.Body = io.NopCloser(strings.NewReader(res.Body))
	httpRes.Proto = res.ProtoVer
	httpRes.ContentLength = res.ContentLength
	httpRes.Header = make(http.Header)

	// TODO: Actually get the propper protocol, but this will do for now
	httpRes.ProtoMajor = 1
	httpRes.ProtoMinor = 1

	// TODO: test this!
	for name, value := range res.Headers {
		for _, subVal := range strings.Split(value, ", ") {
			httpRes.Header.Add(name, subVal)
		}
	}

	return httpRes
}
