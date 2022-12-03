package adapter

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseNoSqlAdapter interface {
	Save(database, collect string, data any) (string, error)
	Find(database, collec string, where *bson.D, op ...*options.FindOptions) (*mongo.Cursor, error)
	FindOne(database, collec string, where *bson.D) *mongo.SingleResult
	Aggregate(database, collec string, pipeline bson.A) (*mongo.Cursor, error)
}
