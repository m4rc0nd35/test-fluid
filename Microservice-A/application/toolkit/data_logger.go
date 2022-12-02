package toolkit

import (
	"encoding/json"

	"github.com/m4rc0nd35/test-fluid/application/adapter"
	"github.com/m4rc0nd35/test-fluid/application/entity"
)

type config struct {
	amqpx adapter.Amqp
}

func NewDataLogger(amqpx adapter.Amqp) *config {
	return &config{amqpx}
}

func (cfg *config) LogQueue(user entity.User) {
	logUser := entity.Log{
		Uuid:       user.Login.Uuid,
		Username:   user.Login.Username,
		Email:      user.Email,
		StatusFlow: user.StatusFlow,
	}

	jsonLog, _ := json.Marshal(logUser)
	cfg.amqpx.SendToQueu("fluid-logs-all", string(jsonLog)) // logs
}
