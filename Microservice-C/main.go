package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/m4rc0nd35/test-fluid/application/repository"
	"github.com/m4rc0nd35/test-fluid/application/toolkit"
	"github.com/m4rc0nd35/test-fluid/domain"
	"github.com/m4rc0nd35/test-fluid/infra/database"
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

	connectionMongo, err := database.NewConnectionMongoDB(
		os.Getenv("MONGODB_HOST"),
		os.Getenv("MONGODB_PORT"),
		os.Getenv("MONGODB_USERNAME"),
		os.Getenv("MONGODB_PASSWORD"),
	)
	toolkit.Error(err)

	logRepo := repository.NewDataLoggerRepository(connectionMongo)

	logDomain := domain.NewDataLogger(rabbitMQ, logRepo)

	// rabbitMQ.ConsumerQueue("fluid-processed-K", 1, func(body string, ch *amqp.Channel, id uint64) {

	// 	// Success
	// 	// ch.Ack(id, false)
	// })

	rabbitMQ.ConsumerQueue("fluid-logs-all", 1, logDomain.DataLogger)

	// Keep alive
	wg.Wait()
}
