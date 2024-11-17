package app

import (
	"epub/internal/mq"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type Epub struct {
	Mq mq.RabbitMq
}

func (epub *Epub) DeclareandBindQueue(name string) (*amqp091.Channel, error) {
	ch, err := epub.Mq.Conn.Channel()
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
