package repository

import (
	"context"
	"encoding/json"
	"errors"
	database_mongo "github.com/Eco-Trading/internal/infra/database"
	"strings"
	"time"

	. "github.com/gobeam/mongo-go-pagination"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GenericQueryRepository[T any] struct {
	MongoDataSource database_mongo.IMongoDatasource
	DatabaseName    string
	TableName       string
}

func (g GenericQueryRepository[T]) FindAllPage(page int, limit int) (
	*PaginatedData,
	[]*T,
	error,
) {
	connect, _ := g.MongoDataSource.Connect()
	defer g.MongoDataSource.Close(connect)
	collection := g.MongoDataSource.DataSource(connect, g.DatabaseName, g.TableName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var mappers []*T
	paginatedData, err := New(collection).Context(ctx).Limit(int64(limit)).Filter(bson.M{}).Page(int64(page)).Decode(&mappers).Find()
	if err != nil {
		return nil, nil, err
	}
	return paginatedData, mappers, nil
}

func (g GenericQueryRepository[T]) FindAll() ([]*T, error) {

	connect, _ := g.MongoDataSource.Connect()
	defer g.MongoDataSource.Close(connect)
	collection := g.MongoDataSource.DataSource(connect, g.DatabaseName, g.TableName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var results []*T
	values, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	if err = values.All(ctx, &results); err != nil {
		return nil, err
	}
	for _, result := range results {
		_, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
	}
	return results, nil
}

func (g GenericQueryRepository[T]) FindByCustomized(key string, column any, orderByColumn string, orderBy string) (*T, error) {
	connect, _ := g.MongoDataSource.Connect()
	defer g.MongoDataSource.Close(connect)
	collection := g.MongoDataSource.DataSource(connect, g.DatabaseName, g.TableName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	valueSort := 1
	if strings.EqualFold("desc", orderBy) {
		valueSort = -1
	}

	filter := bson.D{{Key: key, Value: bson.D{{Key: "$eq", Value: column}}}}
	opts := options.Find().SetSort(bson.D{{orderByColumn, valueSort}})

	var results []*T
	cursor, err := collection.Find(ctx, filter, opts)

	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	for _, result := range results {
		_, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
	}
	if len(results) == 0 {
		return nil, errors.New("not found row")
	}
	return results[0], nil
}

func (g GenericQueryRepository[T]) FindByColumnCustomized(key string, value string) (*T, error) {
	connect, _ := g.MongoDataSource.Connect()
	defer g.MongoDataSource.Close(connect)
	collection := g.MongoDataSource.DataSource(connect, g.DatabaseName, g.TableName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var mapper *T

	filters := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: key, Value: bson.D{{Key: "$eq", Value: value}}}},
			},
		},
	}

	err := collection.FindOne(ctx, filters).Decode(&mapper)

	if err != nil {
		return nil, err
	}
	return mapper, nil
}
