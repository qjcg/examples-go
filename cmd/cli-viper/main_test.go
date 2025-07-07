package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestNewViper(t *testing.T) {
	v := NewViper()

	got := v.GetString("server")
	want := "demo.example.com"
	if got != want {
		t.Fatalf("want: %v, got: %v", want, got)
	}

	gotPort := v.GetInt("port")
	wantPort := 8080
	if gotPort != wantPort {
		t.Fatalf("want: %v, got: %v", wantPort, gotPort)
	}
}

func TestApp_logConfig(t *testing.T) {
	app := &App{
		v: &viper.Viper{},
		config: Config{
			Server: "demo.example.com",
			Port:   uint(8080),
		},
	}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	t.Cleanup(func() {
		log.SetOutput(os.Stderr)
	})

	app.logConfig()

	want := "server: demo.example.com, port: 8080"
	got := buf.String()
	if strings.Contains(got, want) {
		t.Fatalf("want: %v, got %v", want, got)
	}
}
