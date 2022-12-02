package repository

import (
	"fmt"
	"time"

	"github.com/m4rc0nd35/test-fluid/application/adapter"
	"github.com/m4rc0nd35/test-fluid/application/entity"
	"github.com/m4rc0nd35/test-fluid/application/toolkit"
)

type dataLoggerRepository struct {
	noSql adapter.DatabaseNoSqlAdapter
}

func NewDataLoggerRepository(
	noSql adapter.DatabaseNoSqlAdapter,
) *dataLoggerRepository {
	return &dataLoggerRepository{noSql}
}

func (v *dataLoggerRepository) Create(data entity.DataLogger) string {
	now := time.Now()
	data.CreatedAt = now.Format(time.RFC3339)

	// nsec := now.UnixNano()
	ids, err := v.noSql.Save("fluid", "datalogger", data)
	toolkit.Error(err)

	return fmt.Sprint("MySQL: ", 1, " - Mongo: ", ids)
}
