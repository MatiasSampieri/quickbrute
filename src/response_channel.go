package main

import "net/http"

type ResponseChannel struct {
	Channel chan *http.Response
	Open    bool
}

func (ch *ResponseChannel) Close() {
	if ch.Open {
		close(ch.Channel)
		ch.Open = false
	}
}

func (ch *ResponseChannel) Assign(channel chan *http.Response) {
	ch.Channel = channel
	ch.Open = true
}

func (ch *ResponseChannel) Add(res *http.Response) {
	if ch.Open {
		ch.Channel <- res
	}
}