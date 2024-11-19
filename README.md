# signal-api-receiver

# Introduction

### Problem statement

I use the excellent [signal-cli-rest-api][signal-cli-rest-api] by @bbernhard on
my server, and my home-assistant is configured to send me and the Home group
notifications of all kind. Sending messages quickly is crucial for many of my automations,
so I'm running the API [in `json-rpc` mode][exec-mode].

Recently, I wanted to add a way for Home Assistant to receive messages from us
to trigger automations (or stop others), however, when the signal-api is
running in `json-rpc` mode the `/v1/receive` becomes a websocket-only endpoint;
And that is not supported by [Home Assistant's signal_messenger
integration][signal_messenger].

### A possible solution

There are many ways to solve this problem, to name a few:
1. Improve the Home Assistant integration with Signal to function properly with a Websocket.
1. Propose a new endpoint to the `signal-cli-rest-api` that responds to REST.
1. Write a wrapper that consumes the Websocket, stores the messages in memory and expose them over REST.

I wanted this done quickly so I went with the solution with the least friction.
This repository is the implementation of the said wrapper.

[signal-cli-rest-api]: https://github.com/bbernhard/signal-cli-rest-api
[exec-mode]: https://github.com/bbernhard/signal-cli-rest-api?tab=readme-ov-file#execution-modes
[signal_messenger]: https://www.home-assistant.io/integrations/signal_messenger/#sending-messages-to-signal-to-trigger-events
