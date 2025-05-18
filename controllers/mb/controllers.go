package controllers

import "github.com/nats-io/nats.go"

type NatsController interface {
	Connect(conn *nats.Conn)
}
