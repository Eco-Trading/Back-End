package database_mongo

/*
Author: Diego Morais
Data: 10/20/2023
Description: Class responsible for to create connect with database
*/

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
)

type IMongoDatasource interface {
	Connect() (*mongo.Client, error)
	Close(connect *mongo.Client)
	DataSource(client *mongo.Client, database, collection string) *mongo.Collection
	Ping(client *mongo.Client) error
}

var dataSource IMongoDatasource

type mongoDataSource struct {
	clientOptions *options.ClientOptions
}

func New() IMongoDatasource {
	lock := &sync.Mutex{}
	if dataSource == nil {
		lock.Lock()
		defer lock.Unlock()

		dataSource = &mongoDataSource{clientOptions: clientOptions()}
		connect, err := dataSource.Connect()
		if err != nil {
			log.Println(err)
			return nil
		}
		err = dataSource.Ping(connect)
		if err != nil {
			log.Println(err)
			return nil
		}
		defer dataSource.Close(connect)
	}
	return dataSource
}

func clientOptions() *options.ClientOptions {
	uri := os.Getenv("DATABASE_MONGODB_URL")
	return options.Client().ApplyURI(uri)
}

func (c *mongoDataSource) Connect() (*mongo.Client, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, c.clientOptions)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return client, nil
}

func (c *mongoDataSource) Ping(client *mongo.Client) error {
	err := client.Ping(context.Background(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *mongoDataSource) Close(connect *mongo.Client) {
	defer func(client *mongo.Client) {
		err := client.Disconnect(context.Background())
		if err != nil {
			panic(err)
		}
	}(connect)
}

func (c *mongoDataSource) DataSource(client *mongo.Client, database, collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}
