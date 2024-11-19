ARG GO_VERSION=1.23.2
FROM golang:$GO_VERSION AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build ./cmd/signal-api-receiver

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /app/signal-api-receiver /app/signal-api-receiver

WORKDIR /app

EXPOSE 8105

CMD ["/app/signal-api-receiver"]
