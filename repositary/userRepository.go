package repo

import (
	"context"
	"errors"
	"os"
	"real-time-chat-app/database"
	"real-time-chat-app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Declare the collection globally, but initialize it after InitMongoDB is called
var userCollection *mongo.Collection

// Initialize userCollection after the MongoDB connection is established
func InitRepository() {
	// Ensure InitMongoDB is called first
	database.InitMongoDB()

	// Initialize the userCollection
	userCollection = database.GetCollection(os.Getenv("MONGO_TABLE_USER"))
}

// InsertUser inserts a new user into the database
func InsertUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ensure unique email
	filter := bson.M{
		"$or": []bson.M{
			{"email": user.Email},
			{"username": user.Username},
		},
	}

	count, err := userCollection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("email or username already exists")
	}

	// Insert user into the collection
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
