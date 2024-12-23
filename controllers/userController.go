package controllers

import (
	"encoding/json"
	"net/http"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	"real-time-chat-app/services"

	"github.com/gin-gonic/gin"
)

func FetchUserController(c *gin.Context) {

	if c.Request.Method != "GET" {
		logger.LogInfo("FetchUserController :: error GET method required")
		models.ManageResponse(c.Writer, "GET method required", http.StatusMethodNotAllowed, nil, false)
		return
	}

	username := c.DefaultQuery("username", "")
	if len(username) < 1 {
		models.ManageResponse(c.Writer, "please provide the username in query parameter ", http.StatusNotAcceptable, nil, false)
		c.Abort()
		return
	}

	logger.LogInfo("Username fetch from url : " + username)

	userResponse, err := services.UserFetch(username)
	if err != nil {
		logger.LogInfo("FetchUserController :: error  while fetching the user ")
		models.ManageResponse(c.Writer, "error while fetching the user :: "+err.Error(), http.StatusBadRequest, nil, false)
		c.Abort()
		return
	}

	models.ManageResponse(c.Writer, "User fetched succesfully", http.StatusOK, userResponse, true)

}

func UpdateUserAndProfile(c *gin.Context) {

	if c.Request.Method != "POST" {
		logger.LogInfo("UpdateUserAndProfile :: error POST method required")
		models.ManageResponse(c.Writer, "POST method required", http.StatusMethodNotAllowed, nil, false)
		return
	}

	username := c.DefaultQuery("username", "")
	if len(username) < 1 {
		models.ManageResponse(c.Writer, "please provide the username in query parameter ", http.StatusNotAcceptable, nil, false)
		c.Abort()
		return
	}

	var user *models.UpdateUserAndProfile
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&user)
	if err != nil {
		logger.LogInfo("UpdateUserAndProfile :: error in decoding the body" + err.Error())
		models.ManageResponse(c.Writer, "error in decoding the body "+err.Error(), http.StatusBadRequest, nil, false)

		return
	}

	userResponse, err := services.UserAndProfileUpdate(username, user)
	if err != nil {
		logger.LogInfo("UpdateUserAndProfile :: error  while fetching the user ")
		models.ManageResponse(c.Writer, "error while fetching the user :: "+err.Error(), http.StatusBadRequest, nil, false)
		c.Abort()
		return
	}

	models.ManageResponse(c.Writer, "User fetched succesfully", http.StatusOK, userResponse, true)

}
