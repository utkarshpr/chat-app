package repo

import (
	"context"
	"errors"
	"fmt"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveMessage(message *models.Message) error {
	logger.LogInfo("SendMessage repo :: started")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	message.ChatID = message.SenderID + " -> " + message.RecipientID
	_, err := messageCollection.InsertOne(ctx, message)
	if err != nil {
		logger.LogInfo("SendMessage repo :: error " + err.Error())
		return err
	}
	logger.LogInfo("SendMessage repo :: ended")
	return nil
}

func GetMessage(username string, reciever string) ([]*models.GetMessage, error) {
	logger.LogInfo("GetMessage repo :: started")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{
		"sender_id":    username,
		"recipient_id": reciever,
	}

	// var allMessages []bson.M
	// cursor, err := messageCollection.Find(ctx, bson.M{})
	// if err != nil {
	// 	logger.LogError("Error fetching all messages: " + err.Error())
	// 	return nil, errors.New("error fetching messages")
	// }
	// defer cursor.Close(ctx)

	// // Iterate through the cursor and decode each document
	// if err := cursor.All(ctx, &allMessages); err != nil {
	// 	logger.LogError("Error decoding all messages: " + err.Error())
	// 	return nil, errors.New("error decoding messages")
	// }

	// logger.LogInfo(fmt.Sprintf("All Messages: %+v", allMessages))

	logger.LogInfo("GetMessage from" + reciever + " for user " + username)
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "timestamp", Value: 1}})
	fmt.Println(filter)
	cursor, err := messageCollection.Find(ctx, filter, findOptions)
	if err != nil {
		logger.LogError("GetMessage :: error " + err.Error())
		logger.LogError("GetMessage :: error " + err.Error())
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	defer cursor.Close(ctx)
	fmt.Println(cursor)
	// Parse the results into a slice of ContactRequest objects
	var contacts []*models.GetMessage
	for cursor.Next(ctx) {
		var contact models.GetMessage
		if err := cursor.Decode(&contact); err != nil {
			logger.LogError("GetMessage :: error decoding contact: " + err.Error())
			return nil, errors.New("error decoding contact: " + err.Error())
		}
		contacts = append(contacts, &contact)
	}

	fmt.Println(contacts)

	// Check for any errors during cursor iteration
	if err := cursor.Err(); err != nil {
		logger.LogError("GetMessage :: cursor iteration error: " + err.Error())
		return nil, errors.New("error iterating through contacts: " + err.Error())
	}

	logger.LogInfo("GetMessage repo :: ended")
	return contacts, nil
}

func EditMessage(editMessage *models.EditMessage) (*models.Message, error) {
	logger.LogInfo("EditMessage  service :: started ")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{
		"message_id": editMessage.ID,
	}
	fmt.Println(filter)
	var originalmessage *models.Message
	err := messageCollection.FindOne(ctx, filter).Decode(&originalmessage)
	if err != nil {
		logger.LogError("EditMessage repo :: cannot fetch the original message " + err.Error())
		return nil, errors.New("unable to update the message ")
	}
	logger.LogInfo("document fetch ::" + strconv.Itoa(len(filter)))
	originalmessage.Content = editMessage.NewText
	originalmessage.Timestamp = time.Now().Format(time.RFC3339) // Update timestamp manually if needed
	originalmessage.Status = "edited sent"

	// Persist the changes to the database
	update := bson.M{"$set": bson.M{
		"content":   originalmessage.Content,
		"timestamp": originalmessage.Timestamp,
		"status":    originalmessage.Status,
	}}
	_, updateErr := messageCollection.UpdateOne(ctx, filter, update)
	if updateErr != nil {
		logger.LogError("EditMessage repo :: failed to update the message: " + updateErr.Error())
		return nil, errors.New("unable to update the message")
	}
	logger.LogInfo("Edit repo :: content edit successfully")
	return originalmessage, nil
}

func MessageDelete(deleteMessage *models.DeleteMessage) (*models.DeleteMessageResponse, error) {
	logger.LogInfo("MessageDelete service :: started")

	// Prepare the context with a timeout for the database operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println(deleteMessage.ID)
	filter := bson.M{"message_id": deleteMessage.ID}

	deletedResult, err := messageCollection.DeleteOne(ctx, filter)
	if err != nil {
		logger.LogError("MessageDelete service :: error deleting message: " + err.Error())
		return nil, errors.New("error deleting message")
	}

	if deletedResult.DeletedCount == 0 {
		logger.LogError("MessageDelete service :: message not found with ID: " + deleteMessage.ID)
		return nil, errors.New("message not found")
	}

	logger.LogInfo("MessageDelete service :: message deleted successfully, ID: " + deleteMessage.ID)
	messageResponse := &models.DeleteMessageResponse{
		Messsage: "Message deleted Successfully",
	}
	return messageResponse, nil
}
