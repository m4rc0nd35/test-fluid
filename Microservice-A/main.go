package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/m4rc0nd35/test-fluid/application/controller"
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
	ws := controller.NewWebServer()

	// Logger queue
	logs := toolkit.NewDataLogger(rabbitMQ)

	// New leads
	lead := domain.NewLead(rabbitMQ, logs)
	lead.GetLeadApi() // Init
	lead.Start()
	ws.Webserver(lead)

	ws.RunWebServer(":8080")
}
