package services

import (
	"log"
	"net/smtp"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var channel *amqp.Channel

func init() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	channel, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	go CreateHelloWorldQueue(channel)
	// defer channel.Close()
}

func MQConn() *amqp.Channel {
	return channel
}

func CreateHelloWorldQueue(channel *amqp.Channel) {
	q, err := channel.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			auth := smtp.PlainAuth("", "sryoliver@gmail.com", "SRYoliver0603", "smtp.gmail.com")

			// Connect to the server, authenticate, set the sender and recipient,
			// and send the email all in one step.
			to := []string{"rsheng@tilr.com"}
			msg := []byte("To: rsheng@tilr.com\r\n" +
				"Subject: discount Gophers!\r\n" +
				"\r\n" +
				"Hello there.\r\n")
			err := smtp.SendMail("smtp.gmail.com:587", auth, "sryoliver@gmail.com", to, msg)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func RunConsumer() {
	go CreateHelloWorldQueue(channel)
}
