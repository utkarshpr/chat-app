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

	_, err := FetchUserByUsername(contactRequest.ToUserID)
	if err != nil {
		logger.LogError("Contact you were trying to send request doesnot exist")
		return "", errors.New("Contact you were trying to send request doesnot exist " + contactRequest.ToUserID)
	}

	// Define the filter based on from_user_id and to_user_id
	filter := bson.M{
		"from_user_id": contactRequest.FromUserID,
		"to_user_id":   contactRequest.ToUserID,
	}

	var existingRequest models.ContactRequest
	err = contactCollection.FindOne(ctx, filter).Decode(&existingRequest)
	if err != nil && err != mongo.ErrNoDocuments {
		logger.LogError("HandleContactRequest ::error in sending the request" + err.Error())
		return "", errors.New("error finding the contact request" + err.Error())
	}

	if existingRequest.Status == "remove" || existingRequest.Status == "block" {
		if contactRequest.Status != "pending" {
			return "Contact request cannot be sent ,Please send friend request (pending)", nil
		}
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

		mong, err := contactCollection.UpdateOne(ctx, filter, update)

		if err != nil {
			logger.LogError("HandleContactRequest ::error accepting contact request" + err.Error())
			return "", errors.New("error accepting contact request")
		}

		if mong.ModifiedCount > 0 {
			logger.LogInfo("Contact successfully updated.")
			return "Contact request accepted", nil
		} else if mong.UpsertedCount > 0 {
			logger.LogInfo("New contact inserted as part of upsert operation.")
			return "Contact request accepted as a new entry", nil
		} else {
			logger.LogError("HandleContactRequest ::No changes made to the contact request.")
			return "", errors.New("no changes made to the contact request, first send pending request")
		}

	case "rejected":
		// Ensure the request is either pending or accepted before rejecting
		if existingRequest.Status == "rejected" {
			return "contact request already rejected", nil
		}

		if existingRequest.Status == "accepted" {
			return "contact request already accepted can only remove or block", nil
		}

		// Update the request status to "rejected"
		update := bson.M{
			"$set": bson.M{
				"status": "rejected",
			},
		}

		mong, err := contactCollection.UpdateOne(ctx, filter, update)

		if err != nil {
			logger.LogError("HandleContactRequest ::error accepting contact request" + err.Error())
			return "", errors.New("error accepting contact request")
		}

		if mong.ModifiedCount > 0 {
			logger.LogInfo("Contact successfully updated.")
			return "Contact request accepted", nil
		} else if mong.UpsertedCount > 0 {
			logger.LogInfo("New contact inserted as part of upsert operation.")
			return "Contact request accepted as a new entry", nil
		} else {
			logger.LogError("HandleContactRequest ::No changes made to the contact request.")
			return "", errors.New("no changes made to the contact request , first send pending request")
		}

	default:
		return "", errors.New("invalid status provided")
	}
}

func GetAllContactfromUser(username string) ([]*models.ContactRequest, error) {
	logger.LogInfo("GetAllContactfromUser :: starting")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the filter to find contacts where the user is either the "from_user_id" or "to_user_id"
	filter := bson.M{
		"$or": []bson.M{
			{"from_user_id": username},
		},
	}

	// Perform the query to find all matching contacts
	cursor, err := contactCollection.Find(ctx, filter)
	if err != nil {
		logger.LogError("GetAllContactfromUser :: error finding contacts: " + err.Error())
		return nil, errors.New("error finding contacts: " + err.Error())
	}
	defer cursor.Close(ctx)

	// Parse the results into a slice of ContactRequest objects
	var contacts []*models.ContactRequest
	for cursor.Next(ctx) {
		var contact models.ContactRequest
		if err := cursor.Decode(&contact); err != nil {
			logger.LogError("GetAllContactfromUser :: error decoding contact: " + err.Error())
			return nil, errors.New("error decoding contact: " + err.Error())
		}
		contacts = append(contacts, &contact)
	}

	// Check for any errors during cursor iteration
	if err := cursor.Err(); err != nil {
		logger.LogError("GetAllContactfromUser :: cursor iteration error: " + err.Error())
		return nil, errors.New("error iterating through contacts: " + err.Error())
	}

	logger.LogInfo("GetAllContactfromUser :: successfully retrieved contacts")
	return contacts, nil
}

func UpdateContact(contactRequest *models.ContactActionRequest) (string, error) {
	logger.LogInfo("UpdateContact repo:: started")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := FetchUserByUsername(contactRequest.ContactID)
	if err != nil {
		logger.LogError("Contact you were trying to send request doesnot exist")
		return "", errors.New("Contact you were trying to send request doesnot exist " + contactRequest.ContactID)
	}

	// Define the filter based on from_user_id and to_user_id
	filter := bson.M{
		"from_user_id": contactRequest.UserID,
		"to_user_id":   contactRequest.ContactID,
	}

	var existingRequest models.ContactRequest
	err = contactCollection.FindOne(ctx, filter).Decode(&existingRequest)
	if err != nil && err != mongo.ErrNoDocuments {
		logger.LogError("HandleContactRequest ::error in sending the request" + err.Error())
		return "", errors.New("error finding the contact request" + err.Error())
	}
	switch existingRequest.Status {
	case "accepted":
		{

			// Update the request status to "accepted"
			update := bson.M{
				"$set": bson.M{
					"status": contactRequest.Action,
				},
			}

			mong, err := contactCollection.UpdateOne(ctx, filter, update)

			if err != nil {
				logger.LogError("HandleContactRequest ::error accepting contact request" + err.Error())
				return "", errors.New("error accepting contact action")
			}

			if mong.ModifiedCount > 0 {
				logger.LogInfo("Contact successfully updated.")
				return "Contact action accepted", nil
			} else if mong.UpsertedCount > 0 {
				logger.LogInfo("New contact inserted as part of upsert operation.")
				return "Contact action accepted as a new entry", nil
			} else {
				logger.LogError("HandleContactRequest ::No changes made to the contact request.")
				return "", errors.New("no changes made to the contact action, first send pending request")
			}

		}

	}
	logger.LogInfo("UpdateContact repo:: ended")
	return "", errors.New("cannot block or remove cause contact is not connected")
}
