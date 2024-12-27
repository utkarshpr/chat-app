package controllers

import (
	"net/http"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	"real-time-chat-app/security"
	"real-time-chat-app/services"

	"github.com/gin-gonic/gin"
)

// MessageSentController handles sending a new message.
// This endpoint accepts a POST request with a message payload, validates it,
// and ensures the authorized user is the sender.
//
// @Description Sends a new message from the authorized user to the recipient.
// @Tags Messages
// @Accept  json
// @Produce  json
// @Param  requestBody  body  models.Message  true  "Message payload"
// @Success 200  {object}  models.GenericResponse
// @Failure 400  {object}  models.GenericResponse
// @Failure 405  {object}  models.GenericResponse
// @Router /messages/sent [post]
func MessageSentController(c *gin.Context) {
	logger.LogInfo("MessageSentController :: started")

	if c.Request.Method != "POST" {
		logger.LogError("MessageSentController :: Invalid method POST required")
		models.ManageResponse(c.Writer, "Invalid method POST required", http.StatusMethodNotAllowed, nil, false)
		return
	}

	var message models.Message
	if err := c.ShouldBind(&message); err != nil {
		logger.LogError("MessageSentController :: invalid input" + err.Error())
		models.ManageResponse(c.Writer, "Invalid input ", http.StatusBadRequest, nil, false)
		return
	}
	logger.LogInfo("username from message  " + message.SenderID)
	logger.LogInfo("media from message  " + message.MediaURL)
	logger.LogInfo("reciept from message  " + message.RecipientID)

	claims := security.GetClaims(c)
	username := claims["username"].(string)
	if username != message.SenderID {
		logger.LogError("MessageSentController :: Authorize user can only sent the message ")
		models.ManageResponse(c.Writer, "Authorize user can only sent the message  ", http.StatusBadRequest, nil, false)
		return
	}
	//media
	mediaFile, mediaHeader, err := c.Request.FormFile("media")
	if err != nil {
		if err == http.ErrMissingFile {
			logger.LogInfo("No media file uploaded")
		} else {
			logger.LogError("Error retrieving media file: " + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid media file"})
			return
		}
	}

	response, err := services.SendMessage(&message, mediaFile, mediaHeader)
	if err != nil {
		logger.LogError("MessageSentController :: Failed to send message ")
		models.ManageResponse(c.Writer, "Failed to send message"+err.Error(), http.StatusBadRequest, nil, false)
		return
	}
	logger.LogInfo("MessageSentController :: ended")
	models.ManageResponse(c.Writer, "Message sent successfully", http.StatusOK, response, true)
}

// MessageGetAllController retrieves all messages between the authorized user and a specified recipient.
// This endpoint accepts a GET request with the recipient ID as a query parameter.
//
// @Description Fetches all messages exchanged with a specific recipient.
// @Tags Messages
// @Accept  json
// @Produce  json
// @Param  reciever  query  string  true  "Recipient ID"
// @Success 200  {object}  models.GenericResponse
// @Failure 400  {object}  models.GenericResponse
// @Failure 405  {object}  models.GenericResponse
// @Failure 406  {object}  models.GenericResponse
// @Router /messages/get [get]
func MessageGetAllController(c *gin.Context) {
	logger.LogInfo("MessageGetAllController :: started")
	if c.Request.Method != "GET" {
		logger.LogError("MessageGetAllController :: Invalid method GET required")
		models.ManageResponse(c.Writer, "Invalid method GET required", http.StatusMethodNotAllowed, nil, false)
		return
	}

	reciever := c.DefaultQuery("reciever", "")
	if len(reciever) < 1 {
		logger.LogError("please provide the reciever in query parameter ")
		models.ManageResponse(c.Writer, "please provide the reciever in query parameter ", http.StatusNotAcceptable, nil, false)
		c.Abort()
		return
	}

	claims := security.GetClaims(c)
	username := claims["username"].(string)
	response, err := services.GetMessage(username, reciever)
	if err != nil {
		logger.LogError("MessageSentController :: Failed to send message ")
		models.ManageResponse(c.Writer, "Failed to send message"+err.Error(), http.StatusBadRequest, nil, false)
		return
	}

	models.ManageResponse(c.Writer, "Successfully fetched chat for user : "+reciever, http.StatusOK, response, true)

	logger.LogInfo("MessageGetAllController :: ended")
}

func MessageEditController(c *gin.Context) {
	logger.LogInfo("MessageEditController :: started")

	if c.Request.Method != "PATCH" {
		logger.LogError("MessageEditController :: PATCH method is required")
		models.ManageResponse(c.Writer, "PATCH method is required", http.StatusMethodNotAllowed, "", false)
		return
	}

	var editMessage *models.EditMessage
	if err := c.ShouldBindJSON(&editMessage); err != nil {
		logger.LogError("MessageEditController ::uanble to parse the json " + err.Error())
		models.ManageResponse(c.Writer, "uanble to parse the json", http.StatusBadRequest, "", false)
		return
	}

	username := security.GetClaims(c)["username"].(string)
	logger.LogInfo("Username fetch from authorization token " + username)

	if username != editMessage.FromUserID {
		logger.LogError("MessageEditController :: Authorize user can only sent the message ")
		models.ManageResponse(c.Writer, "Authorize user can only sent the message  ", http.StatusBadRequest, nil, false)
		return
	}

	messageResponse, err := services.MessageEdit(editMessage)
	if err != nil {
		logger.LogError("unable to edit message error from service ")
		models.ManageResponse(c.Writer, "unable to edit message error from service ", http.StatusBadRequest, "", false)
		return
	}
	models.ManageResponse(c.Writer, "Successfully edited the message ", http.StatusOK, messageResponse, true)
}

func MessageDeleteController(c *gin.Context) {
	logger.LogInfo("MessageDeleteController :: started")

	if c.Request.Method != "DELETE" {
		logger.LogError("MessageDeleteController :: DELETE method is required")
		models.ManageResponse(c.Writer, "DELETE method is required", http.StatusMethodNotAllowed, "", false)
		return
	}

	var editMessage *models.DeleteMessage
	if err := c.ShouldBindJSON(&editMessage); err != nil {
		logger.LogError("MessageDeleteController ::uanble to parse the json " + err.Error())
		models.ManageResponse(c.Writer, "uanble to parse the json", http.StatusBadRequest, "", false)
		return
	}

	username := security.GetClaims(c)["username"].(string)
	logger.LogInfo("Username fetch from authorization token " + username)

	if username != editMessage.FromUserID {
		logger.LogError("MessageEditController :: Authorize user can only sent the message ")
		models.ManageResponse(c.Writer, "Authorize user can only sent the message  ", http.StatusBadRequest, nil, false)
		return
	}

	_, err := services.MessageDelete(editMessage)
	if err != nil {
		logger.LogError("unable to delete message error from service ")
		models.ManageResponse(c.Writer, "unable to delete message error from service ", http.StatusBadRequest, "", false)
		return
	}
	models.ManageResponse(c.Writer, " MessageDeleteController ::Successfully deleted the message ", http.StatusOK, nil, true)
}
