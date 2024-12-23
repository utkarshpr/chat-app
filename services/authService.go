package services

import (
	"errors"
	"real-time-chat-app/logger"
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

func LoginUser(user *models.LoginUser) (string, string, error) {

	// Delegate to database layer
	logger.LogInfo("LoginUser :: fetching IsLoggedinUserExist")
	token, refreshtoken, err := repo.IsLoggedinUserExist(user)
	if err != nil {
		return "", "", err
	}
	logger.LogInfo("LoginUser ::  JWT token collected .")
	return token, refreshtoken, nil
}
