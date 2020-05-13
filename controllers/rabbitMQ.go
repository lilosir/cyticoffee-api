package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/services/rabbitmq"
)

var messageClient rabbitmq.IMessage

func init() {
	messageClient = &rabbitmq.MessageClient{}
	err := messageClient.ConnectToBroker("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("connect to rabbitmq error", err)
	}
}

// RabbitMQTest for testing
func RabbitMQTest(c *gin.Context) {
	go sendMessage([]byte("hello rabbitmq"))
	// c.JSON(200, "testing")
}

func sendMessage(body []byte) {
	messageClient.PublishToQueue(body, "sendEmail")
}
