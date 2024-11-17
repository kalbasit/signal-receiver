package receiver

import (
	"reflect"
	"strconv"
	"testing"
)

func TestFlush(t *testing.T) {
	t.Run("returns empty list when no messages was found", func(t *testing.T) {
		c := &Client{messages: []Message{}}
		if want, got := []Message{}, c.Flush(); !reflect.DeepEqual(want, got) {
			t.Errorf("want %#v got %#v", want, got)
		}
	})

	t.Run("return the message if only one is there", func(t *testing.T) {
		c := &Client{messages: []Message{Message{Account: "1"}}}

		if want, got := []Message{Message{Account: "1"}}, c.Flush(); !reflect.DeepEqual(want, got) {
			t.Errorf("want %#v got %#v", want, got)
		}
	})

	t.Run("return messages in order", func(t *testing.T) {
		c := &Client{messages: []Message{
			Message{Account: "0"},
			Message{Account: "1"},
			Message{Account: "2"},
		}}

		want := []Message{
			Message{Account: "0"},
			Message{Account: "1"},
			Message{Account: "2"},
		}
		got := c.Flush()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want\n%#v\ngot\n%#v", want, got)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("returns null when no messages was found", func(t *testing.T) {
		c := &Client{messages: []Message{}}
		var want *Message
		if got := c.Pop(); want != got {
			t.Errorf("want %#v got %#v", want, got)
		}
	})

	t.Run("return the message if only one is there", func(t *testing.T) {
		c := &Client{messages: []Message{Message{Account: "1"}}}

		want := Message{Account: "1"}
		got := c.Pop()
		if !reflect.DeepEqual(want, *got) {
			t.Errorf("want\n%#v\ngot\n%#v", want, got)
		}
	})

	t.Run("return messages in order", func(t *testing.T) {
		c := &Client{messages: []Message{
			Message{Account: "0"},
			Message{Account: "1"},
			Message{Account: "2"},
		}}

		for i := range c.messages {
			want := Message{Account: strconv.Itoa(i)}
			got := c.Pop()
			if !reflect.DeepEqual(want, *got) {
				t.Errorf("want\n%#v\ngot\n%#v", want, got)
			}
		}
	})
}
