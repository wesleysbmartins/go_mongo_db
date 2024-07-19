package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoOperations struct{}

type IOperations interface {
	Insert(ctx context.Context, collection string, value interface{}) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, collection string, filters interface{}, value interface{}) error
	Update(ctx context.Context, collection string, filters interface{}, value interface{}) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, collection string, filters interface{}) (*mongo.DeleteResult, error)
}

func (o *MongoOperations) Insert(ctx context.Context, collection string, value interface{}) (*mongo.InsertOneResult, error) {
	result, err := Database.Collection(collection).InsertOne(ctx, value)
	return result, err
}

func (o *MongoOperations) Find(ctx context.Context, collection string, filter interface{}, value interface{}) error {

	cursor, err := Database.Collection(collection).Find(ctx, filter)

	if err != nil {
		return err
	}

	return cursor.All(ctx, value)
}

func (o *MongoOperations) Update(ctx context.Context, collection string, filter interface{}, value interface{}) (*mongo.UpdateResult, error) {
	result, err := Database.Collection(collection).UpdateOne(ctx, filter, value)
	return result, err
}

func (o *MongoOperations) Delete(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error) {
	result, err := Database.Collection(collection).DeleteOne(ctx, filter)
	return result, err
}
