package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func main() {
	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		logrus.Fatal(err)
	}
	defer nc.Close()

	sc, err := stan.Connect("test-cluster", "sub-1", stan.NatsConn(nc))
	if err != nil {
		logrus.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: localhost:4222", err)
	}
	defer sc.Close()

	sub, _ := sc.Subscribe("BruhSub", func(m *stan.Msg) {
		fmt.Printf("GOT: %s\n", string(m.Data))
	})

	sub.Unsubscribe()
}
