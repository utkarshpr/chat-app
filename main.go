package main

import (
	"context"
	"log"
	"real-time-chat-app/database"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	loadEnvVarible()
	// Initialize MongoDB connection
	database.ConnectMongo()

	//TestDatabase()

}

func loadEnvVarible() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func TestDatabase() {
	// Insert a test document into the "users" collection
	usersCollection := database.GetCollection("user")
	testUser := bson.M{
		"username": "test_user1",
		"email":    "test_user@example.com",
		"status":   "active",
		"created":  time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := usersCollection.InsertOne(ctx, testUser)
	if err != nil {
		log.Fatalf("Failed to insert test user: %v", err)
	}
	log.Println("Test user inserted successfully")
}
