package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/m4rc0nd35/test-fluid/application/toolkit"
	"github.com/m4rc0nd35/test-fluid/domain"
	"github.com/m4rc0nd35/test-fluid/infra/service"
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
	logs := domain.NewDataLogger(rabbitMQ)

	// Lead
	leadDomain := domain.NewFlow(rabbitMQ, logs)
	rabbitMQ.ConsumerQueue("fluid-new-H", 1, leadDomain.WorkerNewFlow)

	// Pipeline
	pipeLineDoamin := domain.NewProcessingFlow(rabbitMQ, logs)
	rabbitMQ.ConsumerQueue("fluid-processing-I", 1, pipeLineDoamin.WorkerProcessingFlow)
	rabbitMQ.ConsumerQueue("fluid-recused-J", 1, pipeLineDoamin.Recused)

	// Keep alive
	wg.Wait()
}
