package services

import (
	"context"
	"errors"
	"mime/multipart"
	"real-time-chat-app/config"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	repo "real-time-chat-app/repositary"
	"real-time-chat-app/utils"

	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func SendMessage(message *models.Message, mediaFile multipart.File, mediaHeader *multipart.FileHeader) (*models.Message, error) {
	logger.LogInfo("SendMessage service :: started")
	message.ID = utils.GenerateUUID()
	message.Timestamp = utils.GetCurrentTimestamp()
	message.Status = "sent"

	//cloudinary
	if mediaFile != nil {
		logger.LogInfo("Uploading media to Cloudinary...")
		mediaURL, err := UploadMedia(mediaFile, mediaHeader)
		if err != nil {
			logger.LogError("Failed to upload media to Cloudinary: " + err.Error())
			return nil, errors.New("unable to upload media")
		}
		message.MediaURL = mediaURL
		logger.LogInfo("Media uploaded successfully: " + mediaURL)
	}
	err := repo.SaveMessage(message)
	if err != nil {
		logger.LogError("error in saveing the message ")
		return nil, err
	}
	logger.LogInfo("SendMessage before BroadcastToRecipient" + message.RecipientID)

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
		logger.LogError("error in editing the message ")
		return nil, err
	}
	utils.BroadcastToRecipient(editmessage.ToUserID, editMessageResponse)
	logger.LogInfo("MessageEdit service :: ended ")
	return editMessageResponse, nil
}

func MessageDelete(editmessage *models.DeleteMessage) (*models.DeleteMessageResponse, error) {
	logger.LogInfo("MessageDelete service :: started ")
	editMessageResponse, err := repo.MessageDelete(editmessage)
	if err != nil {
		logger.LogError("error in deleting the message ")
		return nil, err
	}
	utils.BroadcastToRecipientDelete(editmessage.ToUserID, editMessageResponse)
	logger.LogInfo("MessageDelete service :: ended ")
	return editMessageResponse, nil
}

func UploadMedia(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	// Initialize Cloudinary
	cld, err := config.InitCloudinary()
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	// Upload the file
	uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: "message_media", // Folder in your Cloudinary account
	})
	if err != nil {
		return "", err
	}

	// Return the secure URL of the uploaded media
	return uploadResult.SecureURL, nil
}
