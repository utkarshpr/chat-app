package services

import (
	"real-time-chat-app/models"
	repo "real-time-chat-app/repositary"

	"github.com/golang-jwt/jwt"
)

func HandleContactRequest(contactRequest *models.ContactRequest, claims jwt.MapClaims) (string, error) {

	return repo.HandleContactRequest(contactRequest, claims)

}
