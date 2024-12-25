package services

import (
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	repo "real-time-chat-app/repositary"
	"real-time-chat-app/utils"
)

func SendMessage(message *models.Message) (*models.Message, error) {
	logger.LogInfo("SendMessage service :: started")
	message.ID = utils.GenerateUUID()
	message.Timestamp = utils.GetCurrentTimestamp()
	message.Status = "sent"

	err := repo.SaveMessage(message)
	if err != nil {
		return nil, err
	}

	utils.BroadcastToRecipient(message.RecipientID, message)
	logger.LogInfo("SendMessage service :: ended")
	return message, nil
}
