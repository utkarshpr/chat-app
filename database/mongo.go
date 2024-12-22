package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// ConnectMongo establishes a connection to MongoDB.
func ConnectMongo() {
	// Define MongoDB connection URI (update as needed)
	uri := os.Getenv("MONGO_URI")

	// Set up client options
	clientOptions := options.Client().ApplyURI(uri)

	// Create a new MongoDB client
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	// Set a timeout for connection context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect the client to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Select the database
	database := os.Getenv("MONGO_DATABASE")
	DB = client.Database(database)
	log.Println("Connected to MongoDB")
}

// GetCollection returns a MongoDB collection by name.
func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}
