package domain

import (
	"log"

	"github.com/m4rc0nd35/test-fluid/application/adapter"
)

type config struct {
	amqpx adapter.Amqp
}

func NewFlow(amqpx adapter.Amqp) *config {
	return &config{amqpx}
}

func (cfg *config) WorkerNewFlow(body *string) bool {

	log.Println(*body)
	return true
}
