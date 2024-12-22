package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the structure of a user
type User struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Username      string             `json:"username" bson:"username" `
	Email         string             `json:"email" bson:"email" `
	Password      string             `json:"password" bson:"password" `
	AvatarURL     string             `json:"avatar_url" bson:"avatar_url"`
	StatusMessage string             `json:"status_message" bson:"status_message"`
	LastSeen      string             `json:"last_seen" bson:"last_seen"`
	FirstName     string             `json:"first_name" bson:"first_name"`
	LastName      string             `json:"last_name" bson:"last_name"`
	Address       string             `json:"address" bson:"address,omitempty"`
	DateOfBirth   string             `json:"date_of_birth" bson:"date_of_birth"`
}

type UserResponse struct {
	Username    string `json:"username" bson:"username" `
	Email       string `json:"email" bson:"email" `
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	Address     string `json:"address" bson:"address,omitempty"`
	DateOfBirth string `json:"date_of_birth" bson:"date_of_birth"`
}
