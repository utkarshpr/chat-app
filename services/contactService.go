package services

import (
	"errors"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	repo "real-time-chat-app/repositary"

	"github.com/golang-jwt/jwt"
)

func HandleContactRequest(contactRequest *models.ContactRequest, claims jwt.MapClaims) (string, error) {

	return repo.HandleContactRequest(contactRequest, claims)

}

func GetAllContactfromUser(username string) ([]*models.ContactRequest, error) {

	contactResp, err := repo.GetAllContactfromUser(username)
	if err != nil {
		logger.LogError("GetAllContactfromUser :: error in fetching the contact from repo." + err.Error())
		return nil, errors.New(" error in fetching the contact ")

	}

	return contactResp, nil
}
