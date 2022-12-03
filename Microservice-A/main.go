package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/m4rc0nd35/test-fluid/application/controller"
	"github.com/m4rc0nd35/test-fluid/application/repossitory"
	"github.com/m4rc0nd35/test-fluid/application/toolkit"
	"github.com/m4rc0nd35/test-fluid/domain"
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

	// Webserver command
	webserver := controller.NewWebServer()

	// Logger queue
	logs := domain.NewDataLogger(rabbitMQ)

	// New leads
	leadRepo := repossitory.NewLeadApi()
	lead := domain.NewLead(rabbitMQ, leadRepo, logs)
	lead.GetLeadApi() // Init
	lead.Start()
	webserver.Webserver(lead)

	webserver.RunWebServer(":8080")
}
