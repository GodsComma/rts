package mq

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMq struct {
	Conn *amqp.Connection
}

func (mq *RabbitMq) DeclareandBindQueue(name string) (*amqp091.Channel, error) {
	ch, err := mq.Conn.Channel()
	if err != nil {
		fmt.Println("Unable to establiush a channel")
		return &amqp091.Channel{}, nil
	}
	_, err = ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Printf("Unable to declare a new queue with name %s\n", name)
		return &amqp091.Channel{}, err
	}

	return ch, nil
}

func InitMq() RabbitMq {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/test")
	if err != nil {
		fmt.Printf("Unable to dial for a valid mq conn. Error: %s \n", err.Error())
	}
	return RabbitMq{
		Conn: conn,
	}
}
