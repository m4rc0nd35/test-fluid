package main

import (
	"os"
	"runtime"

	"github.com/m4rc0nd35/test-fluid/application/toolkit"
	"github.com/m4rc0nd35/test-fluid/domain"
	"github.com/m4rc0nd35/test-fluid/infra/service"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	runtime.GOMAXPROCS(1)
	ch := make(chan int)

	rabbitMQ, err := service.NewConnectAMQP(
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_AMQP_PORT"),
		os.Getenv("RABBITMQ_USERNAME"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_VHOST"),
	)
	toolkit.Error(err)

	// Logger queue
	logs := toolkit.NewDataLogger(rabbitMQ)

	new := domain.NewFlow(rabbitMQ, logs)

	rabbitMQ.ConsumerQueue("fluid-new-H", 1, func(body string, ch *amqp.Channel, id uint64) {
		err := new.WorkerNewFlow(&body)
		if !err {
			ch.Nack(id, false, true)
			return
		}

		// Success
		ch.Ack(id, false)
	})

	rabbitMQ.ConsumerQueue("fluid-recused-J", 1, func(body string, ch *amqp.Channel, id uint64) {
		new.Recused()
		ch.Ack(id, false)
	})

	<-ch
}
