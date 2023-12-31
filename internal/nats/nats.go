package nats

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"order/internal/entity"
	"order/internal/service"
	"sync"
)

type Nats struct {
	service   *service.Services
	validator *validator.Validate
}

func NewNats(service *service.Services, validator *validator.Validate) *Nats {
	return &Nats{
		service:   service,
		validator: validator,
	}
}

func (n Nats) Connect(clusterId, clientId, natsUrl string) (stan.Conn, error) {
	sc, err := stan.Connect(
		clusterId,
		clientId,
		stan.NatsURL(natsUrl))
	if err != nil {
		logrus.Fatalf("error while connnecting to the nats streaming server: %s", err.Error())
		return sc, err
	}
	logrus.Print("successfully connected to the nats streaming server")

	return sc, nil
}

func (n Nats) Subscribe(wg *sync.WaitGroup, sc stan.Conn, natsSubject string) error {
	defer wg.Done()

	sub, err := sc.Subscribe(natsSubject, func(msg *stan.Msg) {
		message, err := n.UnmarshalTheMessage(msg)
		if err != nil {
			return
		}

		//err = n.service.PutDbOrder(message)
		//if err != nil {
		//	return
		//}
		//n.service.PutCachedOrder(message)
		logrus.Printf("Successfully saved order with uid %s", message.OrderUid)
	})
	if err != nil {
		logrus.Fatalf("error while subscribing to the nats streaming subject: %s", err.Error())
		return err
	}
	for {
		if !sub.IsValid() {
			wg.Done()
			break
		}
	}
	err = sub.Unsubscribe()
	if err != nil {
		logrus.Errorf("error while unsubscribing from the nats streaming subject: %s", err.Error())
		return err
	}
	logrus.Printf("successfully unsubscribed from the nats streaming subject: %s", err.Error())
	return nil
}

func (n Nats) UnmarshalTheMessage(m *stan.Msg) (entity.Order, error) {
	var order entity.Order
	err := json.Unmarshal(m.Data, &order)
	err = n.validator.Struct(&order)
	if err != nil {
		logrus.Errorf("error while unmarshalling message to model : %s", err.Error())
		return order, err
	}
	return order, nil
}
