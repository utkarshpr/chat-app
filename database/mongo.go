package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

// Initialize MongoDB connection
func InitMongoDB() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Replace `your_mongodb_connection_string` with your MongoDB URI
	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping to check connection
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB")
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	// Replace `chat_app_db` with your database name
	database := mongoClient.Database(os.Getenv("MONGO_DATABASE"))
	fmt.Print(database)
	return database.Collection(collectionName)
}
