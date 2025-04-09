// A NATS jetstream consumer that prints received messages.
package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	flagDebug := flag.Bool("d", false, "print debug log messages")
	flagStream := flag.String("s", "ORDERS", "NATS stream name")
	flagURL := flag.String("u", nats.DefaultURL, "NATS URL")
	flagTimeout := flag.Duration("t", 30*time.Second, "timeout")
	flag.Parse()

	if *flagDebug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	nc, err := nats.Connect(*flagURL)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("Connected to NATS", "url", *flagURL)

	js, err := jetstream.New(nc)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Debug("Created JetStream context", "connection", nc)

	ctx, cancel := context.WithTimeout(context.Background(), *flagTimeout)
	defer cancel()

	stream, err := js.Stream(ctx, *flagStream)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1) //nolint:gocritic
	}
	slog.Info("Got JetStream handle", "stream", *flagStream)

	cons, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{})
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("Retrieve consumer handle", "handle", cons)

	cc, err := cons.Consume(func(msg jetstream.Msg) {
		slog.Info("Received jetstream message", "msg", string(msg.Data()))
		err := msg.Ack()
		if err != nil {
			log.Fatalf("error ACKing message: %v", err)
		}
	}, jetstream.ConsumeErrHandler(func(consumeCtx jetstream.ConsumeContext, err error) {
		slog.Error(err.Error())
	}))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer cc.Stop()
	slog.Info("Message handling callback registered")

	defer runtime.Goexit()
}
