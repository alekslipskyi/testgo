package amqp

import (
	"core/logger"
	"encoding/json"
	"github.com/streadway/amqp"
	"os"
	"reflect"
	"time"
)

var log = logger.Logger{Context: "RabbitMQ broker"}

var storedBroker *Broker

type savedMessage struct {
	toQueue string
	body 	interface{}
}

type Broker struct {
	conn    		*amqp.Connection
	channel 		*amqp.Channel
	savedMessages 	[]savedMessage
	isConnected 	bool
}

func Init() Broker {
	if storedBroker != nil {
		return *storedBroker
	} else {
		conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
		log.LogOnError(err, "error from connecting to amqp server ---> ", err)

		ch, err := conn.Channel()
		log.LogOnError(err, "error from connecting to channel --> ", err)
		broker := Broker{conn: conn, channel: ch, isConnected:true}
		storedBroker = &broker

		return broker
	}
}

func (broker *Broker) Close() {
	broker.conn.Close()
	broker.channel.Close()
}

func (broker *Broker) reconnect() {
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))

	log.Warn("Error while trying to connect", err)

	if err == nil {
		ch, err := conn.Channel()
		log.LogOnError(err, "error from connecting to channel --> ", err)

		broker.conn = conn
		broker.channel = ch

		storedBroker = broker

		log.Info("Connection is established")
		broker.isConnected = true
		broker.reSendStoredMessages()
		return
	}

	time.AfterFunc(5 * time.Second, broker.reconnect)
}

func (broker *Broker) reSendStoredMessages() {
	for _, message := range broker.savedMessages {
		broker.SendMessage(message.toQueue, message.body)
	}
}

func (broker *Broker) saveToQueue(toQueue string, body interface{}) {
	broker.savedMessages = append(broker.savedMessages, savedMessage{toQueue, body})
}

func (broker *Broker) SendMessage(toQueue string, body interface{}) {
	if !broker.isConnected {
		broker.saveToQueue(toQueue, body)
		return
	}

	q, err := broker.channel.QueueDeclare(
		toQueue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		broker.isConnected = false
		log.Warn("error from connecting to queue ---> ", err, " reconnection protocol started")
		broker.saveToQueue(toQueue, body)
		broker.reconnect()

		return
	}


	payload, err := json.Marshal(&body)
	log.LogOnError(err, "error from parse expected body to json ---> ", err, "body ---> ", body)

	if string(payload) == "null" {
		switch reflect.TypeOf(body).Kind() {
		case reflect.Slice:
			payload = []byte("[]")
		}
	}

	err = broker.channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType:  "text/application-json",
			Body:         payload,
			DeliveryMode: 2,
		})
	log.LogOnError(err, "error from sending to queue ---> ", err)

	log.Debug("Message was published to queue", toQueue, "with payload", string(payload))
}
