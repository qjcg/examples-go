//go:build integration

package nats

import (
	"context"
	"testing"

	"github.com/nats-io/nats.go"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func newNATSContainer(ctx context.Context) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "nats:2",
		ExposedPorts: []string{"4222/tcp"},
		WaitingFor:   wait.ForLog("Server is ready"),
	}

	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}

func TestWithNATS(t *testing.T) {
	ctx := context.Background()
	natsC, err := newNATSContainer(ctx)
	if err != nil {
		t.Fatalf("failed to create NATS container: %v", err)
	}
	defer func() {
		if err := natsC.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	}()

	endpoint, err := natsC.Endpoint(ctx, "")
	if err != nil {
		t.Fatalf("failed to get NATS endpoint: %v", err)
	}

	nc, err := nats.Connect(endpoint)
	if err != nil {
		t.Fatalf("Error connecting to NATS: %v", err)
	}

	if err := nc.Publish("test", []byte("Hello NATS!")); err != nil {
		t.Fatalf("failed to publish to NATS: %v", err)
	}
}
