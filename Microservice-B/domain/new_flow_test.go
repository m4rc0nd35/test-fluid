package domain

import (
	"testing"

	"github.com/m4rc0nd35/test-fluid/application/toolkit"
	"github.com/m4rc0nd35/test-fluid/infra/service"
	"github.com/stretchr/testify/assert"
)

func TestNewFlow(t *testing.T) {
	// Need rabbitMQ connection up
	rabbitMQ, err := service.NewConnectAMQP(
		"localhost",
		"5672",
		"guest",
		"guest",
		"/",
	)
	toolkit.Error(err)

	// Logger queue
	logs := NewDataLogger(rabbitMQ)

	flow := NewFlow(rabbitMQ, logs)

	assert.True(t, flow.WorkerNewFlow([]byte("{\"test\": 1}")))

}
