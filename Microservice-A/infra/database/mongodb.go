package infra

import (
	"context"
	"fmt"

	"github.com/m4rc0nd35/test-fluid/application/toolkit"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	ctx    context.Context
	Client *mongo.Client
}

func NewConnectionMongoDB(host, port, username, password string) (*MongoDB, error) {
	var db = &MongoDB{}
	var err error

	credential := options.Credential{
		Username: username,
		Password: password,
	}

	clientOpts := options.Client().ApplyURI(fmt.Sprint("mongodb://", host, ":", port)).SetAuth(credential)

	db.Client, err = mongo.Connect(context.Background(), clientOpts)
	toolkit.Error(err)

	return db, nil
}

func (db *MongoDB) Save(database, collect string, data any) (string, error) {

	collection := db.Client.Database(database).Collection(collect)
	result, err := collection.InsertOne(db.ctx, data)
	toolkit.Error(err)

	return fmt.Sprint(result.InsertedID), nil
}
