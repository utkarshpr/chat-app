package models

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the structure of a user
type User struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
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
	Role          Role               `json:"role" bson:"role"`
}

type UserResponse struct {
	Username    string `json:"username" bson:"username" `
	Email       string `json:"email" bson:"email" `
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	Address     string `json:"address" bson:"address,omitempty"`
	DateOfBirth string `json:"date_of_birth" bson:"date_of_birth"`
}

type LoginUser struct {
	Username string `json:"username" bson:"username" `
	Password string `json:"password" bson:"password" `
}

type LoginResponse struct {
	Username     string `json:"username" bson:"username" `
	Token        string `json:"token" bson:"token" `
	RefreshToken string `json:"refreshtoken" bson:"refreshtoken" `
}

type Role string

const (
	Admin  Role = "ADMIN"
	Client Role = "CLIENT"
)

func (r *Role) UnmarshalJSON(data []byte) error {
	var roleStr string
	if err := json.Unmarshal(data, &roleStr); err != nil {
		return err
	}

	switch roleStr {
	case string(Admin), string(Client):
		*r = Role(roleStr)
		return nil
	default:
		return fmt.Errorf("invalid role: %s", roleStr)
	}
}
