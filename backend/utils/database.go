package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// ConnectDatabase establishes connection to MongoDB
func ConnectDatabase() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		return fmt.Errorf("MONGO_URI environment variable is not set")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Test the connection
	if err = client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	DB = client.Database("taskmorph")
	fmt.Println("✅ Connected to MongoDB successfully")
	return nil
}

// GetCollection returns a MongoDB collection
func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}
