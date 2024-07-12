//go:build integration

package pro_bing

import (
	"testing"

	probing "github.com/prometheus-community/pro-bing"
)

func TestICMPPingExampleDotCom(t *testing.T) {
	pinger, err := probing.NewPinger("example.com")
	if err != nil {
		t.Fatalf("Error creating pinger: %v", err)
	}

	pinger.Count = 3
	err = pinger.Run()
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	stats := pinger.Statistics()
	if stats.PacketsSent != 3 || stats.PacketsRecv != 3 {
		t.Errorf("Expected to send and receive 3 packets, but sent: %d, received: %d", stats.PacketsSent, stats.PacketsRecv)
	}
}

func TestICMPPingInvalidHost(t *testing.T) {
	_, err := probing.NewPinger("nonexistent-host")
	if err == nil {
		t.Fatalf("expected error pinging invalid host, got nil")
	}
}
