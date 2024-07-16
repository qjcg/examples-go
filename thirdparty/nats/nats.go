package nats

import (
	"errors"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

func RunEmbeddedServer(inProcess bool) (*nats.Conn, *server.Server, error) {
	serverOpts := &server.Options{
		ServerName:      "embedded_server",
		DontListen:      inProcess,
		JetStream:       true,
		JetStreamDomain: "embedded",
	}

	ns, err := server.NewServer(serverOpts)
	if err != nil {
		return nil, nil, err
	}

	go ns.Start()

	if !ns.ReadyForConnections(5 * time.Second) {
		err := errors.New("timeout waiting for nats server to be ready for connections")
		return nil, nil, err
	}

	clientOpts := []nats.Option{}
	if inProcess {
		clientOpts = append(clientOpts, nats.InProcessServer(ns))
	}

	nc, err := nats.Connect("", clientOpts...)
	if err != nil {
		return nil, nil, err
	}

	return nc, ns, err
}
