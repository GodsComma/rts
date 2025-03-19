package main

import (
	"epub/internal/mq"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type DataStruct struct {
	Id   time.Time
	Data string
}

func fisherYates(data []DataStruct) []DataStruct {
	n := rand.Intn(10)
	if n <= 2 {
		// Buganese
	}
	return data
}

func createFakeMessages(total_size int) []DataStruct {
	new_data := []DataStruct{}
	for i := 1; i < total_size+1; i++ {
		new_data = append(new_data, DataStruct{
			Id:   time.Now(),
			Data: fmt.Sprintf("Test Data %d", i),
		})
	}
	return new_data
}

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
				MessageId:   body,
				Body:        []byte(body),
			},
		)
		msg_count += 1
	}
}

func No_main2() {
	queue_name := "push_pop"
	rabbitMq := mq.InitMq()
	ch, err := rabbitMq.DeclareandBindQueue(queue_name)
	if err != nil {
		fmt.Println("Unable to create a new queue")
	}
	publish_msg(queue_name, 1000000, ch)
}

func main() {
	fmt.Println("DATA", rand.Intn(30))
	fmt.Println("hello world")
	start := time.Now()
	datas := createFakeMessages(300000)
	stop := time.Since(start)
	for _, data := range datas {
		fmt.Printf("%+v\n", data)
	}
	fmt.Printf("Time to print %v\n", stop)

}
