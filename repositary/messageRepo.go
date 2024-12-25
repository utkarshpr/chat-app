package repo

import (
	"context"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	"time"
)

func SaveMessage(message *models.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	logger.LogInfo(messageCollection.Name())
	defer cancel()
	_, err := messageCollection.InsertOne(ctx, message)
	return err
}
