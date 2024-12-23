package validation

import (
	"errors"
	"real-time-chat-app/models"

	"github.com/golang-jwt/jwt"
)

func ValidateAddAndUpdateContact(contactRequest *models.ContactRequest, claim jwt.MapClaims) error {

	if !contactRequest.IsValidStatus() {
		return errors.New("pending, accepted, blocked can only be valid status")
	}

	username := claim["username"].(string)

	if username != contactRequest.FromUserID {
		return errors.New("you can only send a contact request using the logged-in user's username")
	}

	if len(contactRequest.ToUserID) < 1 {
		return errors.New("please provide the username of the contact to whom you want to send a friend request")
	}

	// Ensure "from_user_id" is not the same as "to_user_id"
	if contactRequest.FromUserID == contactRequest.ToUserID {
		return errors.New("you cannot send a request to yourself")

	}

	return nil
}
