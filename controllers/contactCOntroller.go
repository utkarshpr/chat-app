package controllers

import (
	"encoding/json"
	"net/http"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	"real-time-chat-app/security"
	"real-time-chat-app/services"
	"real-time-chat-app/validation"

	"github.com/gin-gonic/gin"
)

// AddAndUpdateContact handles adding or updating a contact request.
//
// @Summary      Adds or updates a contact request between users.
// @Description  This endpoint processes a request to add or update a contact request between two users.
//
//	The request must be made using the POST method, and the payload must include the necessary
//	contact details. Validation is performed on the request body and user claims before processing.
//
// @Tags         Contacts
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer token"
// @Param        body body models.ContactRequest true "Contact request details"
// @Success      200 {object} gin.H{"message": "Contact request processed successfully", "status": true}
// @Failure      405 {object} gin.H{"message": "POST method required", "status": false}
// @Failure      400 {object} gin.H{"message": "Error in decoding or validating the body", "status": false}
// @Failure      500 {object} gin.H{"message": "Error processing the request", "status": false}
// @Router       /contacts/add [post]
func AddAndUpdateCOntact(c *gin.Context) {

	if c.Request.Method != "POST" {
		logger.LogError("AddAndUpdateCOntact :: error POST method required")
		models.ManageResponse(c.Writer, "POST method required", http.StatusMethodNotAllowed, nil, false)
		return
	}

	var contactRequest *models.ContactRequest

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&contactRequest)
	if err != nil {
		logger.LogError("AddAndUpdateCOntact :: error in decoding the body" + err.Error())
		models.ManageResponse(c.Writer, "error in decoding the body "+err.Error(), http.StatusBadRequest, nil, false)

		return
	}

	claims := security.GetClaims(c)
	err = validation.ValidateAddAndUpdateContact(contactRequest, claims)
	if err != nil {
		logger.LogError("AddAndUpdateCOntact :: error in validation the body" + err.Error())
		models.ManageResponse(c.Writer, "error in validating the body "+err.Error(), http.StatusBadRequest, nil, false)
		return
	}

	contactResponse, err := services.HandleContactRequest(contactRequest, claims)

	if err != nil {
		logger.LogError("AddAndUpdateCOntact ::error in sending the request" + err.Error())
		models.ManageResponse(c.Writer, "error in sending the request "+err.Error(), http.StatusBadRequest, nil, false)
		return
	}
	models.ManageResponse(c.Writer, contactResponse, http.StatusOK, nil, true)

}

// GetListofContact retrieves the list of contacts for a specific user.
//
// @Summary      Retrieves the contact list for a user.
// @Description  This endpoint fetches all contacts for a user specified in the query parameter. The user
//               must be authenticated, and the username in the query must match the authorized user's username.
//
// @Tags         Contacts
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer token"
// @Param        username query string true "Username of the user to fetch contacts for"
// @Success      200 {object} gin.H{"message": "Successfully fetched the contact list", "status": true, "data": []models.ContactRequest}
// @Failure      405 {object} gin.H{"message": "GET method required", "status": false}
// @Failure      406 {object} gin.H{"message": "Please provide the username in query parameter", "status": false}
// @Failure      400 {object} gin.H{"message": "Authorized user is not the same as query user", "status": false}
// @Failure      406 {object} gin.H{"message": "Error in fetching the contact list", "status": false}
// @Router       /contacts/get [get]

func GetListofContact(c *gin.Context) {

	if c.Request.Method != "GET" {
		logger.LogError("GetListofContact :: error GET method required")
		models.ManageResponse(c.Writer, "GET method required", http.StatusMethodNotAllowed, nil, false)
		return
	}

	username := c.DefaultQuery("username", "")
	if len(username) < 1 {
		logger.LogError("GetListofContact :: please provide the username in query parameter")
		models.ManageResponse(c.Writer, "please provide the username in query parameter ", http.StatusNotAcceptable, nil, false)
		c.Abort()
		return
	}
	claims := security.GetClaims(c)

	claimUsername := claims["username"].(string)

	if claimUsername != username {
		logger.LogError("GetListofContact :: error in fetching the contact Authorize user is not as same as query user")
		models.ManageResponse(c.Writer, " error in fetching the contact Authorize user is not as same as query user", http.StatusBadRequest, nil, false)
		return
	}

	contactResponse, err := services.GetAllContactfromUser(username)
	if err != nil {
		logger.LogError("GetListofContact :: error in fetching the contact list " + err.Error())
		models.ManageResponse(c.Writer, "error in fetching the contact list ", http.StatusNotAcceptable, nil, false)
		c.Abort()
		return

	}
	logger.LogInfo("successfully fetched the contact of user : " + username)
	models.ManageResponse(c.Writer, "successfully fetched the contact of user : "+username, http.StatusOK, contactResponse, true)

}

func BlockOrRemoveContact(c *gin.Context) {
	if c.Request.Method != "POST" {
		logger.LogError("BlockOrRemoveContact :: method POST is required")
		models.ManageResponse(c.Writer, "method POST is required", http.StatusMethodNotAllowed, nil, false)
		return
	}
	var contactRequest *models.ContactActionRequest

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&contactRequest)
	if err != nil {
		logger.LogError("AddAndUpdateCOntact :: error in decoding the body" + err.Error())
		models.ManageResponse(c.Writer, "error in decoding the body "+err.Error(), http.StatusBadRequest, nil, false)

		return
	}

	err = contactRequest.IsValid()
	if err != nil {
		logger.LogError("BlockOrRemoveContact :: Invalid action " + err.Error())
		models.ManageResponse(c.Writer, " error in payload :: "+err.Error(), http.StatusBadRequest, nil, false)
		return
	}

	claims := security.GetClaims(c)

	claimUsername := claims["username"].(string)

	if claimUsername != contactRequest.UserID {
		logger.LogError("BlockOrRemoveContact :: error in fetching the contact Authorize user is not as same as query user")
		models.ManageResponse(c.Writer, " error in fetching the contact Authorize user is not as same as query user", http.StatusBadRequest, nil, false)
		return
	}

	contactResponse, err := services.UpdateContact(contactRequest)
	if err != nil {
		logger.LogError("BlockOrRemoveContact :: error in updating the contact  " + err.Error())
		models.ManageResponse(c.Writer, err.Error(), http.StatusNotAcceptable, nil, false)
		c.Abort()
		return

	}
	logger.LogInfo("BlockOrRemoveContact ::   " + contactResponse)
	models.ManageResponse(c.Writer, contactResponse, http.StatusOK, nil, true)

}
