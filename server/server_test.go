package server

import "github.com/kalbasit/signal-receiver/receiver"

type mockClient struct {
	msgs []receiver.Message
}

func (mc *mockClient) Pop() *receiver.Message {
	if len(mc.msgs) == 0 {
		return nil
	}

	msg := mc.msgs[0]
	mc.msgs = mc.msgs[1:]

	return &msg
}

func (mc *mockClient) Flush() []receiver.Message {
	msgs := mc.msgs
	mc.msgs = []Message{}
	return msgs
}
