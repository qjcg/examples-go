//go:build integration

package nats

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/testcontainers/testcontainers-go"
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

func TestPubSub(t *testing.T) {
	// FIXME: Remove Skip once issue mentioned below is fixed (e.g. with next Docker Desktop release).
	// https://github.com/testcontainers/testcontainers-go/issues/1359
	// https://github.com/golang/go/issues/61431
	t.Skip()

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

	time.Sleep(2 * time.Second)

	got := strings.TrimSpace(buf.String())
	t.Logf("received %v", got)
	if want != got {
		t.Fatalf("want %v got %v", want, got)
	}
}
