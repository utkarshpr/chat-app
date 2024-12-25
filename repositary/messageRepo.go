package repo

import (
	"context"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	"time"
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
	logger.LogInfo("SendMessage repo :: started")
	return nil
}
