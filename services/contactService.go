package services

import (
	"errors"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	repo "real-time-chat-app/repositary"

	"github.com/golang-jwt/jwt"
)

func HandleContactRequest(contactRequest *models.ContactRequest, claims jwt.MapClaims) (string, error) {
	logger.LogInfo("HandleContactRequest :: starting")
	responseString, err := repo.HandleContactRequest(contactRequest, claims)
	if err != nil {
		logger.LogError("Error from HandleContactRequest repo ")
		return responseString, err
	}
	logger.LogInfo("HandleContactRequest :: exiting")
	return responseString, err
}

func GetAllContactfromUser(username string) ([]*models.ContactRequest, error) {
	logger.LogInfo("GetAllContactfromUser :: starting")
	contactResp, err := repo.GetAllContactfromUser(username)
	if err != nil {
		logger.LogError("GetAllContactfromUser :: error in fetching the contact from repo." + err.Error())
		return nil, errors.New(" error in fetching the contact ")

	}
	logger.LogInfo("GetAllContactfromUser :: ending")
	return contactResp, nil
}

func UpdateContact(contactRequest *models.ContactActionRequest) (string, error) {
	logger.LogInfo("UpdateContact :: starting")
	responseString, err := repo.UpdateContact(contactRequest)
	if err != nil {
		logger.LogError("Error from UpdateContact repo ")
		return responseString, err
	}
	logger.LogInfo("UpdateContact :: ending")
	return responseString, err
}
