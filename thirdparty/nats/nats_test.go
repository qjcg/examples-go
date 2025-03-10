package nats

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

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

func TestMicro_AddService(t *testing.T) {
	nc, ns, err := RunEmbeddedServer(true)
	if err != nil {
		t.Fatal(err)
	}
	defer nc.Close()
	defer ns.Shutdown()

	echoHandler := func(req micro.Request) {
		err := req.Respond(req.Data())
		if err != nil {
			t.Fatalf("error sending echo response: %v", err)
		}
	}

	config := micro.Config{
		Name:    "EchoService",
		Version: "0.1.0",
		Endpoint: &micro.EndpointConfig{
			Subject: "echo",
			Handler: micro.HandlerFunc(echoHandler),
		},
	}

	svc, err := micro.AddService(nc, config)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err := svc.Stop()
		if err != nil {
			t.Fatal(err)
		}
	}()

	testSubject := "greetings"
	want := "foozlefozz"
	var buf bytes.Buffer

	sub, err := nc.Subscribe(testSubject, func(m *nats.Msg) {
		if _, err := fmt.Fprintf(&buf, "%v\n", string(m.Data)); err != nil {
			t.Fatalf("failed to subscribe: %v", err)
		}
	})
	if err != nil {
		t.Fatalf("failed to subscribe: %v", err)
	}

	if err := nc.Publish(testSubject, []byte(want)); err != nil {
		t.Fatalf("failed to publish to NATS: %v", err)
	}

	if err := sub.Drain(); err != nil {
		t.Fatalf("failed to drain subscription: %v", err)
	}
}
