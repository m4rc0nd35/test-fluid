package domain

import (
	"encoding/json"

	"github.com/m4rc0nd35/test-fluid/application/adapter"
	"github.com/m4rc0nd35/test-fluid/application/entity"
)

type dataLogger struct {
	amqpx adapter.Amqp
}

func NewDataLogger(amqpx adapter.Amqp) *dataLogger {
	return &dataLogger{amqpx}
}

func (cfg *dataLogger) LogQueue(user entity.User) bool {
	logUser := entity.Log{
		Uuid:       user.Login.Uuid,
		Username:   user.Login.Username,
		Email:      user.Email,
		StatusFlow: user.StatusFlow,
	}

	jsonLog, err := json.Marshal(logUser)
	if err != nil {
		return false
	}

	cfg.amqpx.SendToQueu("fluid-logs-all", string(jsonLog)) // logs
	return true
}
