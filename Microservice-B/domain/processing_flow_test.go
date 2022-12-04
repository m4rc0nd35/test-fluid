package domain

import (
	"testing"

	"github.com/m4rc0nd35/test-fluid/application/toolkit"
	"github.com/m4rc0nd35/test-fluid/infra/service"
	"github.com/stretchr/testify/assert"
)

func TestProcessingFlow(t *testing.T) {
	// Need rabbitMQ connection up
	rabbitMQ, err := service.NewConnectAMQP(
		"localhost",
		"5672",
		"guest",
		"guest",
		"/",
	)
	toolkit.Error(err)

	// Instance
	newLog := NewDataLogger(rabbitMQ)
	newFlow := NewProcessingFlow(rabbitMQ, newLog)

	assert.True(t, newFlow.WorkerProcessingFlow([]byte("{\"test\": 1}")))
}
