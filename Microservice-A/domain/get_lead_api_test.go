package domain

import (
	"testing"

	"github.com/m4rc0nd35/test-fluid/application/repossitory"
	"github.com/m4rc0nd35/test-fluid/application/toolkit"
	"github.com/m4rc0nd35/test-fluid/infra/service"
	"github.com/stretchr/testify/assert"
)

func TestGetLead(t *testing.T) {
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

	// New leads
	leadRepo := repossitory.NewLeadApi()

	lead := NewLead(rabbitMQ, leadRepo, logs)
	leadCronId := lead.GetLeadApi()

	assert.NotEqual(t, leadCronId, 0)
	assert.Nil(t, err)
	assert.True(t, rabbitMQ.IsConnected)
}
