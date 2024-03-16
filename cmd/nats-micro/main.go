// A minimal NATS micro service.
package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

func main() {
	flagURL := flag.String("u", nats.DefaultURL, "nats URL")
	flag.Parse()

	nc, _ := nats.Connect(*flagURL)

	// request handler
	echoHandler := func(req micro.Request) {
		req.Respond(req.Data())
	}

	_, err := micro.AddService(nc, micro.Config{
		Name:    "EchoService",
		Version: "1.0.0",
		// base handler
		Endpoint: &micro.EndpointConfig{
			Subject: "svc.echo",
			Handler: micro.HandlerFunc(echoHandler),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	runtime.Goexit()
}
