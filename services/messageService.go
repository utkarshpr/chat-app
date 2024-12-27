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
		logger.LogError("error in saveing the message ")
		return nil, err
	}

	utils.BroadcastToRecipient(message.RecipientID, message)
	logger.LogInfo("SendMessage service :: ended")
	return message, nil
}

func GetMessage(username string, reciever string) ([]*models.GetMessage, error) {
	logger.LogInfo("GetMessage service :: started")
	resp, err := repo.GetMessage(username, reciever)
	if err != nil {
		logger.LogError(" GetMessage :: Error in getting message")
		return nil, err
	}

	logger.LogInfo("GetMessage service :: ended")
	return resp, nil

}

func MessageEdit(editmessage *models.EditMessage) (*models.Message, error) {
	logger.LogInfo("MessageEdit service :: started ")
	editMessageResponse, err := repo.EditMessage(editmessage)
	if err != nil {
		logger.LogError("error in saveing the message ")
		return nil, err
	}
	utils.BroadcastToRecipient(editmessage.ToUserID, editMessageResponse)
	logger.LogInfo("MessageEdit service :: ended ")
	return editMessageResponse, nil
}
