package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
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
	hs := httptest.NewServer(&s)
	defer hs.Close()

	t.Run("GET /receive/pop", func(t *testing.T) {
		t.Run("no messages in the queue", func(t *testing.T) {
			mc.msgs = []receiver.Message{}

			resp, err := http.Get(hs.URL + "/receive/pop")
			if err != nil {
				t.Fatalf("expected no error, got: %s", err)
			}
			if want, got := http.StatusNoContent, resp.StatusCode; want != got {
				t.Errorf("want %s, got %s", http.StatusText(want), http.StatusText(got))
			}
		})

		t.Run("one message in the queue", func(t *testing.T) {
			want := receiver.Message{Account: "0"}
			mc.msgs = []receiver.Message{want}

			resp, err := http.Get(hs.URL + "/receive/pop")
			if err != nil {
				t.Fatalf("expected no error, got: %s", err)
			}
			if want, got := http.StatusOK, resp.StatusCode; want != got {
				t.Errorf("want %s, got %s", http.StatusText(want), http.StatusText(got))
			}

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				t.Fatalf("expected no error, got: %s", err)
			}

			var got receiver.Message
			if err := json.Unmarshal(body, &got); err != nil {
				t.Fatalf("expected no error, got: %s", err)
			}

			if !reflect.DeepEqual(want, got) {
				t.Errorf("want\n%#v\ngot\n%#v", want, got)
			}
		})

		t.Run("three messages in the queue", func(t *testing.T) {
			want := []receiver.Message{
				receiver.Message{Account: "0"},
				receiver.Message{Account: "1"},
				receiver.Message{Account: "2"},
			}
			mc.msgs = want

			var got []receiver.Message

			for range want {
				resp, err := http.Get(hs.URL + "/receive/pop")
				if err != nil {
					t.Fatalf("expected no error, got: %s", err)
				}
				if want, got := http.StatusOK, resp.StatusCode; want != got {
					t.Errorf("want %s, got %s", http.StatusText(want), http.StatusText(got))
				}

				body, err := io.ReadAll(resp.Body)
				resp.Body.Close()
				if err != nil {
					t.Fatalf("expected no error, got: %s", err)
				}

				var m receiver.Message
				if err := json.Unmarshal(body, &m); err != nil {
					t.Fatalf("expected no error, got: %s", err)
				}
				got = append(got, m)
			}

			if !reflect.DeepEqual(want, got) {
				t.Errorf("want\n%#v\ngot\n%#v", want, got)
			}
		})
	})

	t.Run("GET /receive/flush", func(t *testing.T) {
		t.Run("no messages in the queue", func(t *testing.T) {
			mc.msgs = []receiver.Message{}

			resp, err := http.Get(hs.URL + "/receive/flush")
			if err != nil {
				t.Fatalf("expected no error, got: %s", err)
			}
			if want, got := http.StatusOK, resp.StatusCode; want != got {
				t.Errorf("want %s, got %s", http.StatusText(want), http.StatusText(got))
			}

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				t.Fatalf("expected no error, got: %s", err)
			}

			var got []receiver.Message
			if err := json.Unmarshal(body, &got); err != nil {
				t.Fatalf("expected no error, got: %s", err)
			}

			if len(got) != 0 {
				t.Errorf("expected an empty response, got\n%#v", got)
			}
		})

		t.Run("one message in the queue", func(t *testing.T) {
			want := []receiver.Message{receiver.Message{Account: "0"}}
			mc.msgs = want

			resp, err := http.Get(hs.URL + "/receive/flush")
			if err != nil {
				t.Fatalf("expected no error, got: %s", err)
			}
			if want, got := http.StatusOK, resp.StatusCode; want != got {
				t.Errorf("want %s, got %s", http.StatusText(want), http.StatusText(got))
			}

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				t.Fatalf("expected no error, got: %s", err)
			}

			var got []receiver.Message
			if err := json.Unmarshal(body, &got); err != nil {
				t.Fatalf("expected no error, got: %s", err)
			}

			if !reflect.DeepEqual(want, got) {
				t.Errorf("want\n%#v\ngot\n%#v", want, got)
			}
		})

		t.Run("three messages in the queue", func(t *testing.T) {
			want := []receiver.Message{
				receiver.Message{Account: "0"},
				receiver.Message{Account: "1"},
				receiver.Message{Account: "2"},
			}
			mc.msgs = want

			resp, err := http.Get(hs.URL + "/receive/flush")
			if err != nil {
				t.Fatalf("expected no error, got: %s", err)
			}
			if want, got := http.StatusOK, resp.StatusCode; want != got {
				t.Errorf("want %s, got %s", http.StatusText(want), http.StatusText(got))
			}

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				t.Fatalf("expected no error, got: %s", err)
			}

			var got []receiver.Message
			if err := json.Unmarshal(body, &got); err != nil {
				t.Fatalf("expected no error, got: %s", err)
			}

			if !reflect.DeepEqual(want, got) {
				t.Errorf("want\n%#v\ngot\n%#v", want, got)
			}
		})
	})

	t.Run("anything else", func(t *testing.T) {
		mc.msgs = []receiver.Message{}

		for _, verb := range []string{"POST", "PUT", "PATCH", "DELETE"} {
			t.Run(verb+" /", func(t *testing.T) {
				r, err := http.NewRequest(verb, hs.URL, nil)
				if err != nil {
					t.Fatalf("expected no error, got: %s", err)
				}
				resp, err := http.DefaultClient.Do(r)
				if err != nil {
					t.Fatalf("expected no error, got: %s", err)
				}
				if want, got := http.StatusForbidden, resp.StatusCode; want != got {
					t.Errorf("want %s, got %s", http.StatusText(want), http.StatusText(got))
				}
			})

			t.Run(verb+" /receive/flush", func(t *testing.T) {
				r, err := http.NewRequest(verb, hs.URL+"/receive/flush", nil)
				if err != nil {
					t.Fatalf("expected no error, got: %s", err)
				}
				resp, err := http.DefaultClient.Do(r)
				if err != nil {
					t.Fatalf("expected no error, got: %s", err)
				}
				if want, got := http.StatusForbidden, resp.StatusCode; want != got {
					t.Errorf("want %s, got %s", http.StatusText(want), http.StatusText(got))
				}
			})

			t.Run(verb+" /receive/pop", func(t *testing.T) {
				r, err := http.NewRequest(verb, hs.URL+"/receive/pop", nil)
				if err != nil {
					t.Fatalf("expected no error, got: %s", err)
				}
				resp, err := http.DefaultClient.Do(r)
				if err != nil {
					t.Fatalf("expected no error, got: %s", err)
				}
				if want, got := http.StatusForbidden, resp.StatusCode; want != got {
					t.Errorf("want %s, got %s", http.StatusText(want), http.StatusText(got))
				}
			})
		}
	})
}
