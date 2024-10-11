package connections

import (
	"context"
	"flashsale-service/internal/config"
	"log"

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

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Println("MongoDB ulanishida xatolik:", err)
	} else {
		log.Println("MongoDB ulanishi muvaffaqiyatli.")
	}

	collection := client.Database("FlashSale").Collection("flashsales")

	return client, collection, nil
}
