package main

import (
	"fmt"
	"net"
	"qdbf/distributed"
)

type NetStatus struct {
	IsHelper bool
	Helpers  []net.Conn
	Parent   net.Conn
}

func (netStat *NetStatus) NetStop() {
	if netStat.IsHelper {
		fmt.Println("STOP: Sending STOP signal to parent...")
		sendAction(netStat.Parent, "STOP")
	} else if len(netStat.Helpers) > 0 {
		fmt.Println("STOP: Sending STOP signal to all helpers...")

		for _, helper := range netStat.Helpers {
			sendAction(helper, "STOP")
		}
	}
}

func (netStat *NetStatus) CheckHelpers() ([]*distributed.SyncMessage, bool) {
	var messages []*distributed.SyncMessage

	for _, helper := range netStat.Helpers {
		msg := checkSync(helper, 50)
		if msg.GetAction() == "STOP" {
			return messages, true
		}

		messages = append(messages, msg)
	}

	return messages, false
}

func (netStat *NetStatus) CheckParentStop() bool {
	msg := checkSync(netStat.Parent, 1000)
	return msg.GetAction() == "STOP"
}