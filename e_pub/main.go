package main

import (
	"epub/internal/mq"
	"fmt"
	"strconv"

	"github.com/rabbitmq/amqp091-go"
)

func publish_msg(queue_name string, count int, ch *amqp091.Channel) {
	msg_count := 1
	for msg_count <= count {
		fmt.Printf("Publishing to queue %s with data %d\n", queue_name, msg_count)
		body := strconv.Itoa(msg_count)
		ch.Publish(
			"",
			queue_name,
			false,
			false,
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
		msg_count += 1
	}
}

func main() {
	queue_name := "push_pop"
	rabbitMq := mq.InitMq()
	ch, err := rabbitMq.DeclareandBindQueue(queue_name)
	if err != nil {
		fmt.Println("Unable to create a new queue")
	}
	publish_msg(queue_name, 1000000, ch)
}
