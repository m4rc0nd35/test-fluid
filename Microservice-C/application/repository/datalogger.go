package repository

import (
	"context"
	"time"

	"github.com/m4rc0nd35/test-fluid/application/adapter"
	"github.com/m4rc0nd35/test-fluid/application/entity"
	"github.com/m4rc0nd35/test-fluid/application/toolkit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dataLoggerRepository struct {
	noSql adapter.DatabaseNoSqlAdapter
}

func NewDataLogger(
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

	return ids
}

func (l *dataLoggerRepository) FindDataLoggerById(uuid string) ([]*entity.DataLogger, error) {
	logs := []*entity.DataLogger{}

	cur, err := l.noSql.Find("fluid", "datalogger", &bson.D{{"uuid", uuid}},
		options.Find().SetProjection(bson.D{
			{"_id", 0},
		}),
		options.Find().SetSort(bson.D{{"createdAt", -1}}),
	)

	for cur.Next(context.Background()) {
		log := entity.DataLogger{}
		cur.Decode(&log)
		logs = append(logs, &log)
	}

	return logs, err
}

func (l *dataLoggerRepository) DataLoggerStats() ([]*entity.Stats, error) {
	stats := []*entity.Stats{}
	filter := bson.A{
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$uuid"},
					{"lastStatusFlow", bson.D{{"$last", "$statusFlow"}}},
				},
			},
		},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$lastStatusFlow"},
					{"count", bson.D{{"$count", bson.D{}}}},
				},
			},
		},
	}

	cur, err := l.noSql.Aggregate("fluid", "datalogger", filter)

	for cur.Next(context.Background()) {
		stat := entity.Stats{}
		cur.Decode(&stat)
		stats = append(stats, &stat)
	}

	if err := cur.Close(context.TODO()); err != nil {
		panic(err)
	}

	return stats, err
}
