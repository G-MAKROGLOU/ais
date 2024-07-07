package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQ is the connection to RabbitMQ
var RabbitMQ *amqp.Connection

// RMQConnect connects to RabbitMQ
func RMQConnect(url string) error {
	conn, err := amqp.Dial(url)

	RabbitMQ = conn

	return err
}
