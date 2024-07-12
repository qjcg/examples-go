package nats

import (
	"log"
	"testing"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

// See https://github.com/synadia-io/rethink_connectivity/blob/main/20-embedding-nats-server/main.go
func TestInProcessConn(t *testing.T) {
	serverOpts := server.Options{
		DontListen: true,
		JetStream:  true,
		StoreDir:   t.TempDir(),
	}

	ns, err := server.NewServer(&serverOpts)
	if err != nil {
		t.Fatal(err)
	}
	go ns.Start()

	if !ns.ReadyForConnections(5 * time.Second) {
		t.Fatal("timeout waiting for nats server to be ready for connections")
	}

	nc, err := nats.Connect(nats.DefaultURL, nats.InProcessServer(ns))
	if err != nil {
		t.Fatal(err)
	}

	nc.Subscribe("hello", func(msg *nats.Msg) {
		log.Println("message received!")
		msg.Respond([]byte("Ahoy there!"))
	})

	ns.WaitForShutdown()
}
