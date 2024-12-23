package repo

import (
	"context"
	"errors"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleContactRequest(contactRequest *models.ContactRequest, claims jwt.MapClaims) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the filter based on from_user_id and to_user_id
	filter := bson.M{
		"from_user_id": contactRequest.FromUserID,
		"to_user_id":   contactRequest.ToUserID,
	}

	var existingRequest models.ContactRequest
	err := contactCollection.FindOne(ctx, filter).Decode(&existingRequest)
	if err != nil && err != mongo.ErrNoDocuments {
		logger.LogError("HandleContactRequest ::error in sending the request" + err.Error())
		return "", errors.New("error finding the contact request" + err.Error())
	}

	switch contactRequest.Status {
	case "pending":
		if err == mongo.ErrNoDocuments {
			// If no existing request is found, create a new one with "pending" status
			contactRequest.Status = "pending" // Ensure status is set to "pending"
			_, err := contactCollection.InsertOne(ctx, contactRequest)
			if err != nil {
				logger.LogError("HandleContactRequest ::error sending contact request" + err.Error())
				return "", errors.New("error sending contact request")
			}
			return "Contact request sent successfully", nil
		}

		// If the existing request is in an "accepted" or "rejected" state, update to "pending"
		if existingRequest.Status == "accepted" {
			return "already accepted, cannot change to pending", nil
		}

		if existingRequest.Status == "pending" {
			return "contact request already in pending state", nil
		}

		// Update existing request to "pending"
		update := bson.M{
			"$set": bson.M{
				"status": "pending",
			},
		}

		_, err = contactCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			logger.LogError("HandleContactRequest ::error updating contact request to pending" + err.Error())
			return "", errors.New("error updating contact request to pending")
		}

		return "Contact request updated to pending", nil

	case "accepted":
		// Ensure the request is in pending state before accepting
		if existingRequest.Status == "rejected" {
			return "already rejected, cannot accept", nil
		}

		if existingRequest.Status == "accepted" {
			return "contact request already accepted", nil
		}

		// Update the request status to "accepted"
		update := bson.M{
			"$set": bson.M{
				"status": "accepted",
			},
		}

		_, err = contactCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			logger.LogError("HandleContactRequest ::error accepting contact request" + err.Error())
			return "", errors.New("error accepting contact request")
		}

		return "Contact request accepted", nil

	case "rejected":
		// Ensure the request is either pending or accepted before rejecting
		if existingRequest.Status == "rejected" {
			return "contact request already rejected", nil
		}

		// Update the request status to "rejected"
		update := bson.M{
			"$set": bson.M{
				"status": "rejected",
			},
		}

		_, err = contactCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			logger.LogInfo("HandleContactRequest ::error rejecting contact request" + err.Error())
			return "", errors.New("error rejecting contact request")
		}

		return "Contact request rejected", nil

	default:
		return "", errors.New("invalid status provided")
	}
}
