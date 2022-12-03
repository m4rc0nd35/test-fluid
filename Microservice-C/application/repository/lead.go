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

type leadRepository struct {
	noSql adapter.DatabaseNoSqlAdapter
}

func NewLead(noSql adapter.DatabaseNoSqlAdapter) *leadRepository {
	return &leadRepository{noSql}
}

func (l *leadRepository) Create(data entity.User) string {
	now := time.Now()
	data.CreatedAt = now.Format(time.RFC3339)

	// nsec := now.UnixNano()
	ids, err := l.noSql.Save("fluid", "datalake", data)
	toolkit.Error(err)

	return ids
}

func (l *leadRepository) FindOneLead(uuid string) (*entity.User, error) {
	lead := entity.User{}
	filter := bson.D{{"login.uuid", uuid}}
	cur := l.noSql.FindOne("fluid", "datalake", &filter).Decode(&lead)
	return &lead, cur
}

func (l *leadRepository) FindAllLead() ([]*entity.User, error) {
	leads := []*entity.User{}

	cur, err := l.noSql.Find("fluid", "datalake", &bson.D{},
		options.Find().SetProjection(bson.D{
			{"_id", 0},
		}),
		options.Find().SetSort(bson.D{{"createdAt", -1}}),
	)

	for cur.Next(context.Background()) {
		lead := entity.User{}
		cur.Decode(&lead)
		// fmt.Println(lead)
		leads = append(leads, &lead)
	}

	return leads, err
}
