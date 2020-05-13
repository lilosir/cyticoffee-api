package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

const (
	QueueName = "sendEmail"
)

var messageClient IMessage

func InitSendEmailQueue() {
	messageClient = &MessageClient{}
	err := messageClient.ConnectToBroker("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("connect to rabbitmq error", err)
	}

	messageClient.SubscribeToQueue(QueueName, sendEmail)
	if err != nil {
		fmt.Println("Failed to comsuer the msg", err)
	}
}

func sendEmail(delivery amqp.Delivery) {
	fmt.Println("sending email....", string(delivery.Body))
}
