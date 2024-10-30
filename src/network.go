package main

import (
	"fmt"
	"net"
	"qdbf/distributed"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"
)


func checkSync(conn net.Conn, timeout time.Duration) *distributed.SyncMessage {
	if timeout == 0 {
		conn.SetReadDeadline(time.Now().Add(time.Millisecond * timeout))
	} else {
		conn.SetReadDeadline(time.Time{})
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
	_, err = conn.Write(resData)

	return err
}


func sendAction(conn net.Conn, action string) error {
	return sendSync(conn, &distributed.SyncMessage{Action: action})
}


func waitForMainInstance(port int) net.Conn {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("Could not start listening on port:", port)
		return nil
	}
	
	fmt.Println("Waiting for main instance to communicate...")
	
	var conn net.Conn
	for {
		conn, err = listener.Accept()
		if err != nil {
			fmt.Println("Could not establish connection!")
		}

		break
	}

	fmt.Println("Main instance conneted! Wating for HELLO...")

	data := make([]byte, 4096)
	count, err := conn.Read(data)
	if err != nil {
		panic(err)
	}

	msg := distributed.SyncMessage{}
	err = proto.Unmarshal(data[:count], &msg)
	if err != nil {
		fmt.Println("Bad request!")
		return nil
	}

	if msg.GetAction() == "HELLO" {
		fmt.Println("Recieved HELLO from main instance!")
		sendAction(conn, "ACK")
	}

	fmt.Println("Main instance conneted succesfully!")

	return conn
}


func waitForRemoteStart(conn net.Conn) {
	fmt.Println("Witing for main instance to send START signal...")

	//msg := checkSync(conn, 0)
}


func askForHelp(addr string, port int) net.Conn {
	fmt.Println("Attempting connection with helper:", addr)

	conn, err := net.Dial("tcp", addr+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("Could not connect to helper")
		return nil
	}

	fmt.Println("Connected to helper, sending HELLO")
	sendAction(conn, "HELLO")

	res := checkSync(conn, 5000)
	if res.GetAction() != "ACK" {
		fmt.Println("Did not recieved ACK from helper, aborting")
		conn.Close()
		return nil
	}

	fmt.Println("Connected succesfully to helper!")

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