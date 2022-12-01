package main

import (
	"fmt"
	"os"

	"github.com/m4rc0nd35/test-fluid/application/controller"
	"github.com/m4rc0nd35/test-fluid/application/toolkit"
	"github.com/m4rc0nd35/test-fluid/domain"
	"github.com/m4rc0nd35/test-fluid/infra/service"
)

func main() {
	rabbitMQ, err := service.NewConnectAMQP(
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_AMQP_PORT"),
		os.Getenv("RABBITMQ_USERNAME"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_VHOST"),
	)
	toolkit.Error(err)

	lead := domain.NewLead(rabbitMQ)
	lead.GetLeadApi() // Init
	lead.Start()

	// Webserver command
	ws := controller.NewWebServer()

	ws.Webserver(lead)
	ws.RunWebServer(fmt.Sprint(":", os.Getenv("WS_PORT")))
}
