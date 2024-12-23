package services

import (
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	repo "real-time-chat-app/repositary"
)

func UserFetch(username string) (*models.UserResponse, error) {

	userResponse, err := repo.UserFetchFromDB(username)
	if err != nil {
		logger.LogInfo("UserFetch :: error  while fetching the user ")
		return nil, err
	}
	return userResponse, nil

}

func UserAndProfileUpdate(username string, user *models.UpdateUserAndProfile) (*models.UserResponse, error) {
	userResponse, err := repo.UserAndProfileUpdate(username, user)
	if err != nil {
		logger.LogInfo("UserAndProfileUpdate :: error  while fetching the user ")
		return nil, err
	}
	return userResponse, nil
}

func DeleteUser(username string) error {
	err := repo.DeleteUser(username)
	if err != nil {
		logger.LogInfo("DeleteUser :: error  while fetching the user ")
		return err
	}
	return nil
}
