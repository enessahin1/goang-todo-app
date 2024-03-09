package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func ConnectDB() *mongo.Database {
	errorENV := godotenv.Load()

	if errorENV != nil {
		panic("Failed to load env file")
	}

	uri := os.Getenv("DATABASE_URL")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("test")

}

func DisconnectDB() {
	err := collection.Database().Client().Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
