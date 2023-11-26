package main

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func main() {
	// Connect to NATS
	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		logrus.Fatal(err)
	}
	defer nc.Close()

	sc, err := stan.Connect("test-cluster", "me", stan.NatsConn(nc))
	if err != nil {
		logrus.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: localhost:4222", err)
	}
	defer sc.Close()

	for i := 1; ; i++ {
		sc.Publish("Bruh", []byte("bruh "+strconv.Itoa(i)))
		time.Sleep(2 * time.Second)
	}
}
