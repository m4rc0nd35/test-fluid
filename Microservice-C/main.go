package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/m4rc0nd35/test-fluid/application/controller"
	"github.com/m4rc0nd35/test-fluid/application/repository"
	"github.com/m4rc0nd35/test-fluid/application/toolkit"
	"github.com/m4rc0nd35/test-fluid/domain"
	"github.com/m4rc0nd35/test-fluid/infra/database"
	"github.com/m4rc0nd35/test-fluid/infra/service"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("v1.0.0")

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

	webserver := controller.NewWebServer()

	// DataLogger
	logRepo := repository.NewDataLogger(connectionMongo)
	logDomain := domain.NewDataLogger(logRepo)
	rabbitMQ.ConsumerQueue("fluid-logs-all", 1, logDomain.DataLogger)
	webserver.DataLoggerOneWS(logDomain)
	webserver.DataLoggerStatsWS(logDomain)

	// Lead
	leadRepo := repository.NewLead(connectionMongo)
	leadDomain := domain.NewLead(leadRepo)
	rabbitMQ.ConsumerQueue("fluid-processed-K", 1, leadDomain.Lead)
	webserver.LeadOneWS(leadDomain) // Endpoint lead by uuid
	webserver.LeadAllWS(leadDomain) // Endpoint leads all

	webserver.RunWebServer(":8081")
}
