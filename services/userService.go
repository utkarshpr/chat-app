package services

import (
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	repo "real-time-chat-app/repositary"
)

func UserFetch(username string) (*models.UserResponse, error) {

	userResponse, err := repo.UserFetchFromDB(username)
	if err != nil {
		logger.LogInfo("FetchUserController :: error  while fetching the user ")
		return nil, err
	}
	return userResponse, nil

}
