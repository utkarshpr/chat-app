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
	logger.LogInfo("CreateUser service :: started")
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.LogError("error in GenerateFromPassword")
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	// Delegate to database layer
	err = repo.InsertUser(user)
	if err != nil {
		return err
	}
	logger.LogInfo("CreateUser service :: ended")
	return nil
}

func LoginUser(user *models.LoginUser) (string, string, error) {

	// Delegate to database layer
	logger.LogInfo("LoginUser service :: fetching IsLoggedinUserExist")
	token, refreshtoken, err := repo.IsLoggedinUserExist(user)
	if err != nil {
		return "", "", err
	}
	logger.LogInfo("LoginUser service ::  JWT token collected .")
	return token, refreshtoken, nil
}

func LogoutUser(username string) error {
	logger.LogInfo(" LogoutUser service :: started")
	err := repo.LogoutUser(username)
	if err != nil {
		logger.LogInfo("LogoutUser :: ailed to logout the user")
		return errors.New("failed to logout the user")
	}
	logger.LogInfo("LogoutUser service :: ended")
	return nil
}
