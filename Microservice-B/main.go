package main

import (
	"os"

	"github.com/m4rc0nd35/test-fluid/application/toolkit"
	"github.com/m4rc0nd35/test-fluid/domain"
	"github.com/m4rc0nd35/test-fluid/infra/service"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch := make(chan int)
	rabbitMQ, err := service.NewConnectAMQP(
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_AMQP_PORT"),
		os.Getenv("RABBITMQ_USERNAME"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_VHOST"),
	)
	toolkit.Error(err)

	new := domain.NewFlow(rabbitMQ)

	rabbitMQ.ConsumerQueue("new-lead", 1, func(body string, ch *amqp.Channel, id uint64) {
		success := new.WorkerNewFlow(&body)
		if success {
			ch.Ack(id, false)
		}
	})

	// rabbitMQ.ConsumerQueue(os.Getenv("QUEUE_NEW"), 1, func(body string, ch *amqp.Channel, id uint64) {

	// })

	<-ch
}
