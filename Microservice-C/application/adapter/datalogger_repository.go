package adapter

import "github.com/m4rc0nd35/test-fluid/application/entity"

type DataLoggerRepository interface {
	Create(data entity.DataLogger) string
	FindDataLoggerById(uuid string) ([]*entity.DataLogger, error)
	DataLoggerStats() ([]*entity.Stats, error)
}

type DataLoggerDomain interface {
	FindDataLoggerById(uuid string) []*entity.DataLogger
	DataLoggerStats() []*entity.Stats
}
