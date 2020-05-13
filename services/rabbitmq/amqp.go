package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

// IMessage is the interface of rabbitmq
type IMessage interface {
	ConnectToBroker(connstr string) error
	PublishToQueue(body []byte, queueName string) error
	SubscribeToQueue(queueName string, handlerFunc func(amqp.Delivery)) error
	Close()
}

// MessageClient rabbitmq client
type MessageClient struct {
	conn *amqp.Connection
}

// failOnError handle error
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(msg)
	}
}

// Close implement close rabbitmq
func (m *MessageClient) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

// ConnectToBroker implements connect to rabbitmq
func (m *MessageClient) ConnectToBroker(connstr string) error {
	// mqconnstr = "amqp://guest:guest@localhost:5672/"
	var err error
	m.conn, err = amqp.Dial(connstr)
	failOnError(err, "Failed to connect to RabbitMQ")
	return err
}

// PublishToQueue publish message to queue
func (m *MessageClient) PublishToQueue(body []byte, queueName string) error {
	if m.conn == nil {
		panic("Please connect ot rabbitmq first!")
	}
	ch, err := m.conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // queue name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")

	return err
}

// SubscribeToQueue will listen for the messages
func (m *MessageClient) SubscribeToQueue(queueName string, handlerFunc func(amqp.Delivery)) error {
	if m.conn == nil {
		panic("Please connect ot rabbitmq first!")
	}

	ch, err := m.conn.Channel()
	failOnError(err, "Failed to open a channel")
	// defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // queue name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go consumingMessages(msgs, handlerFunc)

	return err
}

func consumingMessages(msgs <-chan amqp.Delivery, handlerFunc func(amqp.Delivery)) {
	go func() {
		for d := range msgs {
			// log.Printf("Received a message: %s", d.Body)
			handlerFunc(d)
		}
	}()
}
