package services

import (
	"errors"
	"real-time-chat-app/models"
	repo "real-time-chat-app/repositary"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser handles the business logic for creating a user
func CreateUser(user *models.User) error {

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	// Delegate to database layer
	err = repo.InsertUser(user)
	if err != nil {
		return err
	}

	return nil
}
