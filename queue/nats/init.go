package nats

import (
	"fmt"
	"github.com/cuwand/pondasi/logger"
	"github.com/nats-io/nats.go"
	"time"
)

var natsConnectionClient *nats.Conn

type Nats struct {
	client *nats.Conn
	logger logger.Logger
}

func InitConnection(natsAddress, username, password string, logger logger.Logger) Nats {
	opt := nats.Options{
		Url:            natsAddress,
		AllowReconnect: true,
		MaxReconnect:   10,
		ReconnectWait:  5 * time.Second,
		Timeout:        1 * time.Second,
		User:           username,
		Password:       password,
	}

	nc, err := opt.Connect()

	if err != nil {
		logger.Error(fmt.Sprintf("failed connect to NATS server : %s", natsAddress))
		panic(fmt.Sprintf("failed connect to NATS server : %s", natsAddress))
	}

	logger.Info("NATS Connected")

	return Nats{
		client: nc,
		logger: logger,
	}
}

func (r Nats) GetNatsClient() *nats.Conn {
	return r.client
}

func (r Nats) Publish(subject string, body []byte) error {
	r.logger.Info(fmt.Sprintf("Publish with subject: %s - %s", subject, string(body)))

	return r.client.Publish(subject, body)
}

func (r Nats) RegisterHandlerConsumerGroup(subject, groupId string, workerCount int, cb nats.MsgHandler) error {
	r.logger.Info(fmt.Sprintf("Registering NATS consumer, subject: %s", subject))

	_, err := r.client.QueueSubscribe(subject, groupId, cb)

	return err
}
