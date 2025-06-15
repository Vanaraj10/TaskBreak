package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}

	// Test the connection
	if err = client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	DB = client.Database("taskmorph")
	fmt.Println("Connected to MongoDB")
}

func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}
