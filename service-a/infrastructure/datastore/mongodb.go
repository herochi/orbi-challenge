package datastore

import (
	"context"
	"log"

	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB() (*mongo.Database, error) {
	ENV := viper.GetString("ENV")
	log.Println("Env: ", ENV)
	uri := viper.GetString("MONGO.URI")
	log.Println("MONGO.URI: ", uri)
	if ENV == "DEV" {
		uri = viper.GetString("MONGO.URI_DEV")
	}
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	log.Println("MONGO.DATABASE: ", viper.GetString("MONGO.DATABASE"))
	return client.Database(viper.GetString("MONGO.DATABASE")), nil
}
