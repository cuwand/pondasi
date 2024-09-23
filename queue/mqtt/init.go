package mqtt

import (
	"fmt"
	"github.com/cuwand/pondasi/logger"
	"time"
)
import mqtt "github.com/eclipse/paho.mqtt.golang"

type Mqtt struct {
	client mqtt.Client
	logger logger.Logger
}

func InitConnection(mqttAddress, clientID, username, password string, logger logger.Logger) Mqtt {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", mqttAddress))
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetClientID(clientID)

	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(2 * time.Second)
	opts.OnConnect = func(client mqtt.Client) {

	}

	client := mqtt.NewClient(opts)

	token := client.Connect()

	for !token.WaitTimeout(3 * time.Second) {
	}

	if err := token.Error(); err != nil {
		panic(err)
	}

	return Mqtt{
		client: client,
		logger: logger,
	}
}

func (r Mqtt) GetMqttClient() mqtt.Client {
	return r.client
}

func (r Mqtt) Publish(topic string, body []byte) error {
	r.logger.Info(fmt.Sprintf("Publish with topic: %s - %s", topic, string(body)))

	return r.client.Publish(topic, 0, false, body).Error()
}

func (r Mqtt) RegisterHandlerConsumerGroup(topic, clientId string, workerCount int, cb mqtt.MessageHandler) error {
	r.logger.Info(fmt.Sprintf("Registering MQTT consumer, topic: %s", topic))

	fmt.Println("START")

	//for {
	if token := r.client.Subscribe(topic, 2, cb); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	//}

	fmt.Println("END")

	return nil
	//return r.client.Subscribe(topic, 2, cb).Error()
}
