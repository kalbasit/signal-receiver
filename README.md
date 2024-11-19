# signal-api-receiver

## Introduction

### Problem statement

I use the excellent [signal-cli-rest-api][signal-cli-rest-api] by @bbernhard on
my server, and my home-assistant is configured to send me and the Home group
notifications of all kinds. Sending messages quickly is crucial for many of my
automations, so I'm running the API [in `json-rpc` mode][exec-mode].

Recently, I wanted to add a way for Home Assistant to receive messages from us
to trigger automations (or stop others). However, when the signal-api is
running in `json-rpc` mode, the `/v1/receive` endpoint becomes websocket-only.
This is not supported by [Home Assistant's signal_messenger
integration][signal_messenger], which relies on REST API calls.

### Solution

This project, `signal-api-receiver`, provides a solution by creating a
lightweight wrapper that:

* Consumes the websocket stream from the `/v1/receive` endpoint.
* Stores received messages in memory.
* Exposes a REST API for retrieving those messages.

This approach allows Home Assistant to easily receive Signal messages and
trigger automations without requiring modifications to the existing
`signal-cli-rest-api` or the Home Assistant integration.

### Alternative Solutions

While developing `signal-api-receiver` solved my immediate need, there were
other potential approaches to this problem:

1. Improve the Home Assistant integration with Signal to function properly with a Websocket.
2. Propose a new endpoint to the `signal-cli-rest-api` that responds to REST.

These alternatives might be more comprehensive solutions in the long term, but
creating the wrapper provided a more immediate and focused solution for my
specific use case.

[signal-cli-rest-api]: https://github.com/bbernhard/signal-cli-rest-api
[exec-mode]: https://github.com/bbernhard/signal-cli-rest-api?tab=readme-ov-file#execution-modes
[signal_messenger]: https://www.home-assistant.io/integrations/signal_messenger/#sending-messages-to-signal-to-trigger-events
