package database

import (
	"context"
	"fmt"

	"github.com/m4rc0nd35/test-fluid/application/toolkit"
	"go.mongodb.org/mongo-driver/bson"
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

func (db *MongoDB) Find(database, collec string, where *bson.D, op ...*options.FindOptions) (*mongo.Cursor, error) {

	// Find data
	coll := db.Client.Database(database).Collection(collec)
	cursor, err := coll.Find(db.ctx, where, op...)
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

func (db *MongoDB) FindOne(database, collec string, where *bson.D) *mongo.SingleResult {

	// Find data
	coll := db.Client.Database(database).Collection(collec)
	result := coll.FindOne(db.ctx, where)

	return result
}

func (db *MongoDB) Aggregate(database, collec string, pipeline bson.A) (*mongo.Cursor, error) {

	// Aggregations
	coll := db.Client.Database(database).Collection(collec)
	cursor, err := coll.Aggregate(db.ctx, pipeline)
	if err != nil {
		fmt.Println("Aggregate", err)
		return nil, err
	}
	return cursor, nil
}
