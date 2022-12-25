package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("RabbitMQ in Golang!")

	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to RabbitMQ instance.")
}
