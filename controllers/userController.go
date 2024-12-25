package controllers

import (
	"encoding/json"
	"net/http"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	"real-time-chat-app/security"
	"real-time-chat-app/services"

	"github.com/gin-gonic/gin"
)

// FetchUserController handles HTTP GET requests to fetch user details based on a username.
//
// @Summary Fetch user details
// @Description Handles requests to retrieve user details using a username provided in the query parameters.
// @Tags Users
// @Accept json
// @Produce json
// @Param username query string true "Username of the user to fetch"
// @Success 200 {object} models.UserResponse "User details fetched successfully"
// @Failure 400 {object} models.ErrorResponse "Error while fetching user details"
// @Failure 405 {object} models.ErrorResponse "GET method required"
// @Failure 406 {object} models.ErrorResponse "Username not provided in query parameter"
// @Router /fetchUser [get]
func FetchUserController(c *gin.Context) {
	logger.LogInfo("FetchUserController :: started")
	if c.Request.Method != "GET" {
		logger.LogError("FetchUserController :: error GET method required")
		models.ManageResponse(c.Writer, "GET method required", http.StatusMethodNotAllowed, nil, false)
		return
	}

	username := c.DefaultQuery("username", "")
	if len(username) < 1 {
		logger.LogError("please provide the username in query parameter")
		models.ManageResponse(c.Writer, "please provide the username in query parameter ", http.StatusNotAcceptable, nil, false)
		c.Abort()
		return
	}

	logger.LogInfo("Username fetch from url : " + username)

	userResponse, err := services.UserFetch(username)
	if err != nil {
		logger.LogError("FetchUserController :: error  while fetching the user ")
		models.ManageResponse(c.Writer, "error while fetching the user :: "+err.Error(), http.StatusBadRequest, nil, false)
		c.Abort()
		return
	}
	logger.LogInfo("FetchUserController ::ended")
	models.ManageResponse(c.Writer, "User fetched succesfully", http.StatusOK, userResponse, true)

}

// UpdateUserAndProfile handles HTTP POST requests to update a user's profile details.
//
// @Summary Update user and profile details
// @Description Updates the details of a user based on the provided username and JSON body payload.
// @Tags Users
// @Accept json
// @Produce json
// @Param username query string true "Username of the user to update"
// @Param body body models.UpdateUserAndProfile true "JSON payload containing updated user details"
// @Success 200 {object} models.UserResponse "User details updated successfully"
// @Failure 400 {object} models.ErrorResponse "Error in decoding the request body"
// @Failure 405 {object} models.ErrorResponse "POST method required"
// @Failure 406 {object} models.ErrorResponse "Username not provided in query parameter"
// @Router /updateUser [post]
func UpdateUserAndProfile(c *gin.Context) {

	logger.LogInfo("UpdateUserAndProfile :: started")
	if c.Request.Method != "POST" {
		logger.LogError("UpdateUserAndProfile :: error POST method required")
		models.ManageResponse(c.Writer, "POST method required", http.StatusMethodNotAllowed, nil, false)
		return
	}

	username := c.DefaultQuery("username", "")
	if len(username) < 1 {
		logger.LogError("please provide the username in query parameter ")
		models.ManageResponse(c.Writer, "please provide the username in query parameter ", http.StatusNotAcceptable, nil, false)
		c.Abort()
		return
	}
	logger.LogInfo("UpdateUserAndProfile :: username in query param " + username)
	var user *models.UpdateUserAndProfile
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&user)
	if err != nil {
		logger.LogError("UpdateUserAndProfile :: error in decoding the body" + err.Error())
		models.ManageResponse(c.Writer, "error in decoding the body "+err.Error(), http.StatusBadRequest, nil, false)

		return
	}

	userResponse, err := services.UserAndProfileUpdate(username, user)
	if err != nil {
		logger.LogError("UpdateUserAndProfile :: error  while updating the user ")
		models.ManageResponse(c.Writer, "error while updating the user :: "+err.Error(), http.StatusBadRequest, nil, false)
		c.Abort()
		return
	}
	logger.LogInfo("UpdateUserAndProfile :: ended ")
	models.ManageResponse(c.Writer, "User fetched succesfully", http.StatusOK, userResponse, true)

}

func DeleteUserController(c *gin.Context) {
	logger.LogInfo("DeleteUserController :: started")
	if c.Request.Method != "DELETE" {
		logger.LogError("DeleteUserController :: error DELETE method required")
		models.ManageResponse(c.Writer, "DELETE method required", http.StatusMethodNotAllowed, nil, false)
		return
	}

	username := c.DefaultQuery("username", "")
	if len(username) < 1 {
		logger.LogError("please provide the username in query parameter ")
		models.ManageResponse(c.Writer, "please provide the username in query parameter ", http.StatusNotAcceptable, nil, false)
		c.Abort()
		return
	}
	logger.LogInfo("DeleteUserController :: username is " + username)
	claims := security.GetClaims(c)
	role := claims["role"].(string)

	if role == "ADMIN" {

		err := services.DeleteUser(username)
		if err != nil {
			logger.LogError("DeleteUserController :: error  while deleting the user ")
			models.ManageResponse(c.Writer, "error while deleting the user :: "+err.Error(), http.StatusBadRequest, nil, false)
			c.Abort()
			return
		}
		logger.LogInfo("DeleteUserController :: ended ")
		models.ManageResponse(c.Writer, "User Deleted succesfully", http.StatusOK, nil, true)
	} else {
		logger.LogError("DeleteUserController :: ended NON ADMIN")
		models.ManageResponse(c.Writer, "User cannot be deleted :: ADMIN role required ", http.StatusBadRequest, nil, false)
	}
}
