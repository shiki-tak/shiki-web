package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBHandler struct {
	MongoCli *mongo.Client
}

func NewMongoDBHandler(url string) (*MongoDBHandler, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return &MongoDBHandler{MongoCli: client}, nil
}

func (handler *MongoDBHandler) Put(key []byte, value []byte, args ...interface{}) error {
	return nil
}

func (handler *MongoDBHandler) Get(key string, args ...interface{}) ([]byte, error) {
	return nil, nil
}

func (handler *MongoDBHandler) Scan() {
}
