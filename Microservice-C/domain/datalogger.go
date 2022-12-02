package domain

import (
	"encoding/json"

	"github.com/m4rc0nd35/test-fluid/application/adapter"
	"github.com/m4rc0nd35/test-fluid/application/entity"
)

type config struct {
	amqpx   adapter.Amqp
	logRepo adapter.DataLoggerRepository
}

func NewDataLogger(amqpx adapter.Amqp, logRepo adapter.DataLoggerRepository) *config {
	return &config{amqpx, logRepo}
}

func (cfg *config) LogQueue(user entity.User) {
	logUser := entity.DataLogger{
		Uuid:       user.Login.Uuid,
		Username:   user.Login.Username,
		Email:      user.Email,
		StatusFlow: user.StatusFlow,
	}

	jsonLog, _ := json.Marshal(logUser)
	cfg.amqpx.SendToQueu("fluid-logs-all", string(jsonLog)) // logs
}

func (d *config) DataLogger(body []byte) bool {
	log := entity.DataLogger{}

	// json to struct
	if err := json.Unmarshal(body, &log); err != nil {
		return false
	}

	// insert log
	d.logRepo.Create(log)
	return true
}
