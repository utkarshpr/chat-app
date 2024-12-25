package controllers

import (
	"net/http"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	"real-time-chat-app/security"
	"real-time-chat-app/services"

	"github.com/gin-gonic/gin"
)

func MessageSentController(c *gin.Context) {
	logger.LogInfo("MessageSentController :: started")

	if c.Request.Method != "POST" {
		logger.LogError("MessageSentController :: Invalid method POST required")
		models.ManageResponse(c.Writer, "Invalid method POST required", http.StatusMethodNotAllowed, nil, false)
		return
	}

	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		logger.LogError("MessageSentController :: invalid input" + err.Error())
		models.ManageResponse(c.Writer, "Invalid input ", http.StatusBadRequest, nil, false)
		return
	}
	claims := security.GetClaims(c)
	username := claims["username"].(string)
	if username != message.SenderID {
		logger.LogError("MessageSentController :: Authorize user can only sent the message ")
		models.ManageResponse(c.Writer, "Authorize user can only sent the message  ", http.StatusBadRequest, nil, false)
		return
	}

	response, err := services.SendMessage(&message)
	if err != nil {
		logger.LogError("MessageSentController :: Failed to send message ")
		models.ManageResponse(c.Writer, "Failed to send message"+err.Error(), http.StatusBadRequest, nil, false)
		return
	}
	logger.LogInfo("MessageSentController :: ended")
	models.ManageResponse(c.Writer, "Message sent successfully", http.StatusOK, response, true)
}

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
