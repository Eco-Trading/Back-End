package repository

/*
Author: Diego Morais
Data: 10/20/2023
Description: Class generic Command
*/

import (
	"context"
	"github.com/Eco-Trading/internal/infra/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GenericCommandRepository[T any] struct {
	MongoDataSource database_mongo.IMongoDatasource
	DatabaseName    string
	TableName       string
}

func (g GenericCommandRepository[T]) Save(mapper T) (string, error) {
	connect, _ := g.MongoDataSource.Connect()
	defer g.MongoDataSource.Close(connect)
	collection := g.MongoDataSource.DataSource(connect, g.DatabaseName, g.TableName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := collection.InsertOne(ctx, &mapper)
	if err != nil {
		return "", err
	}
	return response.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (g GenericCommandRepository[T]) Update(id string, mapper *T) error {
	connect, _ := g.MongoDataSource.Connect()
	defer g.MongoDataSource.Close(connect)
	collection := g.MongoDataSource.DataSource(connect, g.DatabaseName, g.TableName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": bson.M{"$eq": id}}
	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": mapper})
	if err != nil {
		return err
	}
	return nil

}

func (g GenericCommandRepository[T]) Delete(id string) error {
	connect, _ := g.MongoDataSource.Connect()
	defer g.MongoDataSource.Close(connect)
	collection := g.MongoDataSource.DataSource(connect, g.DatabaseName, g.TableName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
