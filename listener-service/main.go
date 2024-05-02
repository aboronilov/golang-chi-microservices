package main

import (
	"fmt"
	"listener/event"
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// connect to RMQ
	rabbitConn, err := connect()
	if err != nil {
		panic(err)
	}
	defer rabbitConn.Close()

	// start listening to messages
	fmt.Println("Listening for RMQ queue and consuming messages...")

	// create a consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	// watch the queue and consume events
	topics := []string{"log.INFO", "log.WARNING", "log.ERROR"}
	err = consumer.Listen(topics)
	if err != nil {
		panic(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost:5672")
		if err != nil {
			fmt.Println("RabbitMQ is not ready yet")
			counts++
		} else {
			fmt.Println("Connected to RabbitMQ")
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Backing off...")
		time.Sleep(backOff)
		continue
	}
	return connection, nil
}
