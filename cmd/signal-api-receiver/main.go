package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/kalbasit/signal-api-receiver/receiver"
	"github.com/kalbasit/signal-api-receiver/server"
)

var addr string
var signalApiURL string
var signalAccount string
var repeatLastMessage bool

func init() {
	flag.StringVar(&addr, "addr", ":8105", "The address to listen and serve on")
	flag.StringVar(&signalApiURL, "signal-api-url", "", "The URL of the Signal api including the scheme. e.g wss://signal-api.example.com")
	flag.StringVar(&signalAccount, "signal-account", "", "The account number for signal")
	flag.BoolVar(&repeatLastMessage, "repeat-last-message", false, "Repeat the last message if there are no new messages (applies to /receive/pop)")
}

func main() {
	flag.Parse()

	uri, err := url.Parse(signalApiURL)
	if err != nil {
		log.Printf("error parsing the url %q: %s", signalApiURL, err)
		return
	}
	if uri.Scheme == "" {
		log.Printf("the given url %q does not contain a scheme", uri)
		return
	}
	if uri.Host == "" {
		log.Printf("the given url %q does not contain a host", uri)
		return
	}

	uri.Path = fmt.Sprintf("/v1/receive/%s", signalAccount)
	log.Printf("the fully qualified URL for signal-api was computed as %q", uri.String())

	sarc, err := receiver.New(uri)
	if err != nil {
		log.Printf("error creating a new websocket connetion: %s", err)
		panic(err)
	}

	srv := server.New(sarc, repeatLastMessage)

	log.Print("Starting HTTP server on :8105")
	http.ListenAndServe(addr, srv)
}
