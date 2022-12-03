package domain

import (
	"encoding/json"

	"github.com/m4rc0nd35/test-fluid/application/adapter"
	"github.com/m4rc0nd35/test-fluid/application/entity"
)

type config struct {
	logRepo adapter.DataLoggerRepository
}

func NewDataLogger(logRepo adapter.DataLoggerRepository) *config {
	return &config{logRepo}
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

func (d *config) FindDataLoggerById(uuid string) []*entity.DataLogger {
	log, _ := d.logRepo.FindDataLoggerById(uuid)
	return log
}

func (d *config) DataLoggerStats() []*entity.Stats {
	stats, _ := d.logRepo.DataLoggerStats()
	return stats
}
