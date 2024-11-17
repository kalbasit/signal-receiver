package server

import (
	"testing"

	"github.com/kalbasit/signal-receiver/receiver"
)

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
	mc.msgs = []receiver.Message{}
	return msgs
}

func TestServeHTTP(t *testing.T) {
	mc := &mockClient{msgs: []receiver.Message{}}
	s := Server{sarc: mc}

	t.Run("GET /receive/pop", func(t *testing.T) {

	})

	t.Run("GET /receive/flush", func(t *testing.T) {

	})

	t.Run("anything else", func(t *testing.T) {
		mc.msgs = []receiver.Message{}

		for _, verb := range []string{"POST", "PUT", "PATCH", "DELETE"} {

		}
	})
}
