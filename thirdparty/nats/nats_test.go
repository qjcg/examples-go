package nats

import (
	"testing"

	"github.com/nats-io/nats.go"
)

// See https://github.com/synadia-io/rethink_connectivity/blob/main/20-embedding-nats-server/main.go
func TestInProcessConn(t *testing.T) {
	nc, ns, err := RunEmbeddedServer(true)
	if err != nil {
		t.Fatal(err)
	}
	defer nc.Close()
	defer ns.Shutdown()

	if !nc.IsConnected() {
		t.Fatal("client is not connected to nats-server")
	}
}

func TestPubSub(t *testing.T) {
	nc, ns, err := RunEmbeddedServer(true)
	if err != nil {
		t.Fatal(err)
	}
	defer nc.Close()
	defer ns.Shutdown()

	_, err = nc.Subscribe("greetings", func(msg *nats.Msg) {
		t.Log("message received!")
		err := msg.Respond([]byte("Ahoy there!"))
		if err != nil {
			t.Log(err)
		}
	})
	if err != nil {
		t.Fatal(err)
	}

	err = nc.Publish("greetings", []byte("hello there!"))
	if err != nil {
		t.Fatal(err)
	}
}
