//go:build integration

package nats

import (
	"testing"

	"github.com/nats-io/nats.go"
)

func TestPublish(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := nc.Publish("foo", []byte("Hello World")); err != nil {
		t.Fatal(err)
	}
}
