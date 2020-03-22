package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/services"
	"github.com/streadway/amqp"
)

//MQ for testing rabbitmq
func MQ(c *gin.Context) {
	message := c.PostForm("message")

	channel := services.MQConn()
	q, err := channel.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments

	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	// body := "Hello World!"
	err = channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}
}
