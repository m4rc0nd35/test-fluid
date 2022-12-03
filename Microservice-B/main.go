package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/m4rc0nd35/test-fluid/application/toolkit"
	"github.com/m4rc0nd35/test-fluid/domain"
	"github.com/m4rc0nd35/test-fluid/infra/service"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("v1.0.0")
	var wg sync.WaitGroup
	wg.Add(1)

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

	newLead := domain.NewFlow(rabbitMQ, logs)
	processingLead := domain.NewProcessingFlow(rabbitMQ, logs)

	rabbitMQ.ConsumerQueue("fluid-new-H", 1, func(body string, ch *amqp.Channel, id uint64) {
		err := newLead.WorkerNewFlow(&body)
		if !err {
			ch.Nack(id, false, true)
			return
		}

		// Success
		ch.Ack(id, false)
	})

	// Queue I
	rabbitMQ.ConsumerQueue("fluid-processing-I", 1, func(body string, ch *amqp.Channel, id uint64) {
		err := processingLead.WorkerProcessingFlow(&body)
		if !err {
			ch.Nack(id, false, true)
			return
		}

		// Success
		ch.Ack(id, false)
	})

	// Queue J
	rabbitMQ.ConsumerQueue("fluid-recused-J", 1, func(body string, ch *amqp.Channel, id uint64) {
		processingLead.Recused()
		ch.Ack(id, false)
	})

	// Keep alive
	wg.Wait()
}
