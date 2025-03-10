// A minimal NATS micro service.
package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

func echoHandler(req micro.Request) {
	err := req.Respond(req.Data())
	if err != nil {
		log.Println(err)
	}
}

func main() {
	flagURL := flag.String("u", nats.DefaultURL, "nats URL")
	flag.Parse()

	nc, err := nats.Connect(*flagURL)
	if err != nil {
		log.Fatal(err)
	}

	_, err = micro.AddService(nc, micro.Config{
		Name:    "EchoService",
		Version: "0.1.0",
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
