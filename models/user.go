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
	Profile       Profile            `json:"profile" bson:"profile"`
}

type Profile struct {
	Bio                 string            `json:"bio" bson:"bio"`
	IsProfilePublic     bool              `json:"is_profile_public" bson:"is_profile_public"`
	CoverPhotoURL       string            `json:"cover_photo_url" bson:"cover_photo_url"`
	ProfileCompleteness int               `json:"profile_completeness" bson:"profile_completeness"`
	SocialLinks         map[string]string `json:"social_links" bson:"social_links"`
	Interests           []string          `json:"interests" bson:"interests"`
	ContactPreferences  map[string]bool   `json:"contact_preferences" bson:"contact_preferences"`
	Occupation          string            `json:"occupation" bson:"occupation"`
	Education           string            `json:"education" bson:"education"`
	Achievements        []string          `json:"achievements" bson:"achievements"`
	Gender              string            `json:"gender" bson:"gender"`
	PhoneNumber         string            `json:"phone_number,omitempty" bson:"phone_number,omitempty"`
}

type UpdateUserAndProfile struct {
	Password      string  `json:"password" bson:"password" `
	AvatarURL     string  `json:"avatar_url" bson:"avatar_url"`
	StatusMessage string  `json:"status_message" bson:"status_message"`
	LastSeen      string  `json:"last_seen" bson:"last_seen"`
	FirstName     string  `json:"first_name" bson:"first_name"`
	LastName      string  `json:"last_name" bson:"last_name"`
	Address       string  `json:"address" bson:"address,omitempty"`
	DateOfBirth   string  `json:"date_of_birth" bson:"date_of_birth"`
	Role          Role    `json:"role" bson:"role"`
	Profile       Profile `json:"profile" bson:"profile"`
}
type UserResponse struct {
	Username    string  `json:"username" bson:"username" `
	Email       string  `json:"email" bson:"email" `
	FirstName   string  `json:"first_name" bson:"first_name"`
	LastName    string  `json:"last_name" bson:"last_name"`
	Address     string  `json:"address" bson:"address,omitempty"`
	DateOfBirth string  `json:"date_of_birth" bson:"date_of_birth"`
	Profile     Profile `json:"profile" bson:"profile"`
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
