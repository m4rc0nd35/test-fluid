package domain

import (
	"testing"

	"github.com/m4rc0nd35/test-fluid/application/entity"
	"github.com/m4rc0nd35/test-fluid/application/toolkit"
	"github.com/m4rc0nd35/test-fluid/infra/service"
	"github.com/stretchr/testify/assert"
)

func TestDataLoggerFlow(t *testing.T) {
	// Need rabbitMQ connection up
	rabbitMQ, err := service.NewConnectAMQP(
		"localhost",
		"5672",
		"guest",
		"guest",
		"/",
	)
	toolkit.Error(err)

	user := entity.User{}
	user.Login.Uuid = "983de3ce-93b9-4cc5-8ef2-5c8d6cca59ac"
	user.Login.Username = "fluid"
	user.Email = "test@fluid.io"
	user.StatusFlow = "processing"

	// Instance
	newLog := NewDataLogger(rabbitMQ)

	assert.True(t, newLog.LogQueue(user))
}
