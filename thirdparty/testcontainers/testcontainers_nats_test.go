//go:build integration

package testcontainers

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"github.com/testcontainers/testcontainers-go"
	tcNATS "github.com/testcontainers/testcontainers-go/modules/nats"
	"github.com/testcontainers/testcontainers-go/wait"
)

func newNATSContainer(ctx context.Context) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "nats:2",
		ExposedPorts: []string{"4222/tcp"},
		WaitingFor: wait.ForAll(
			wait.ForLog("Server is ready"),
			wait.ForExposedPort().WithStartupTimeout(180*time.Second),
			wait.ForListeningPort("4222/tcp").WithStartupTimeout(10*time.Second),
		),
	}

	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}

func TestPubSub_testcontainer(t *testing.T) {
	ctx := context.Background()
	natsContainer, err := tcNATS.Run(ctx, "nats:2.10-alpine")
	if err != nil {
		t.Fatalf("failed to create NATS container: %s", err)
	}
	defer func() {
		if err := natsContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	}()

	endpoint, err := natsContainer.Endpoint(ctx, "")
	if err != nil {
		t.Fatalf("failed to get NATS endpoint: %v", err)
	}

	nc, err := nats.Connect(endpoint)
	if err != nil {
		t.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()

	const (
		testSubject string = "test"
		want        string = "Hello, NATS!"
	)
	var buf bytes.Buffer

	sub, err := nc.Subscribe(testSubject, func(m *nats.Msg) {
		if _, err := fmt.Fprintf(&buf, "%v\n", string(m.Data)); err != nil {
			t.Fatalf("failed to subscribe: %v", err)
		}
	})

	if err := nc.Publish(testSubject, []byte(want)); err != nil {
		t.Fatalf("failed to publish to NATS: %v", err)
	}

	if err := sub.Drain(); err != nil {
		t.Fatalf("failed to drain subscription: %v", err)
	}

	time.Sleep(3 * time.Second)

	got := strings.TrimSpace(buf.String())
	t.Logf("received %v", got)
	if want != got {
		t.Fatalf("want %v got %v", want, got)
	}
}

func ExampleMicroService() {
	nc, err := nats.Connect("127.0.0.1:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	echoHandler := func(req micro.Request) {
		req.Respond(req.Data())
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
		log.Fatal(err)
	}
	defer svc.Stop()

	runtime.Goexit()
}
