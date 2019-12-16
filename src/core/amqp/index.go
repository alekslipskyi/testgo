package amqp

import (
	"core/logger"
	"encoding/json"
	"github.com/streadway/amqp"
	"os"
	"reflect"
)

var log = logger.Logger{Context: "RabbitMQ broker"}

type Broker struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func Init() Broker {
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	log.LogOnError(err, "error from connecting to amqp server ---> ", err)

	ch, err := conn.Channel()
	log.LogOnError(err, "error from connecting to channel --> ", err)

	return Broker{conn: conn, channel: ch}
}

func (broker *Broker) Close() {
	broker.conn.Close()
	broker.channel.Close()
}

func (broker *Broker) SendMessage(toQueue string, body interface{}) {
	q, err := broker.channel.QueueDeclare(
		toQueue,
		true,
		false,
		false,
		false,
		nil,
	)
	log.LogOnError(err, "error from connecting to queue ---> ", err)

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
