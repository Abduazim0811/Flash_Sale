package connections

import (
	"context"
	"order_service/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongodb() (*mongo.Client, *mongo.Collection, error) {
	c := config.Configuration()
	clientOptions := options.Client().ApplyURI(c.Mongo.Url)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, nil, err
	}

	collection := client.Database("Orders").Collection("order")

	return client, collection, nil
}
