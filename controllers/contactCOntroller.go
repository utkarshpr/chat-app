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

func AddAndUpdateCOntact(c *gin.Context) {

	if c.Request.Method != "POST" {
		logger.LogInfo("AddAndUpdateCOntact :: error POST method required")
		models.ManageResponse(c.Writer, "POST method required", http.StatusMethodNotAllowed, nil, false)
		return
	}

	var contactRequest *models.ContactRequest

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&contactRequest)
	if err != nil {
		logger.LogInfo("AddAndUpdateCOntact :: error in decoding the body" + err.Error())
		models.ManageResponse(c.Writer, "error in decoding the body "+err.Error(), http.StatusBadRequest, nil, false)

		return
	}

	claims := security.GetClaims(c)
	err = validation.ValidateAddAndUpdateContact(contactRequest, claims)
	if err != nil {
		logger.LogInfo("AddAndUpdateCOntact :: error in validation the body" + err.Error())
		models.ManageResponse(c.Writer, "error in validating the body "+err.Error(), http.StatusBadRequest, nil, false)
		return
	}

	contactResponse, err := services.HandleContactRequest(contactRequest, claims)

	if err != nil {
		logger.LogInfo("AddAndUpdateCOntact ::error in sending the request" + err.Error())
		models.ManageResponse(c.Writer, "error in sending the request "+err.Error(), http.StatusBadRequest, nil, false)
		return
	}
	models.ManageResponse(c.Writer, contactResponse, http.StatusOK, nil, true)

}
