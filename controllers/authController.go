package controllers

import (
	"encoding/json"
	"net/http"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	"real-time-chat-app/services"
	"real-time-chat-app/validation"
)

// SignUp godoc
// @Summary Sign up a new user
// @Description User registers by providing username, email, and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User registration details"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Bad Request"
// @Router /signup [post]
func SignUpController(w http.ResponseWriter, r *http.Request) {

	// Only allow POST method
	if r.Method != "POST" {
		logger.LogInfo("SignUpController :: error POST method required")
		models.ManageResponse(w, "POST method required", http.StatusMethodNotAllowed, nil, false)
		return
	}

	var user models.User

	// Decode the JSON body into the User struct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		logger.LogInfo("SignUpController :: error in decoding the body")
		models.ManageResponse(w, "error in decoding the body", http.StatusBadRequest, nil, false)

		return
	}

	// validation
	err = validation.SignUpUserValidation(&user)
	if err != nil {
		logger.LogInfo("SignUpController :: error in validation  " + err.Error())
		models.ManageResponse(w, "Error : "+err.Error(), http.StatusBadRequest, nil, false)
		return
	}

	// Call the service to handle sign-up logic
	err = services.CreateUser(&user)
	if err != nil {
		logger.LogInfo("SignUpController :: error in service call " + err.Error())
		models.ManageResponse(w, "Unable to create the User "+err.Error(), http.StatusBadRequest, nil, false)
		return
	}

	// // Return success response
	responseModel := &models.UserResponse{
		Username:    user.Username,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Address:     user.Address,
		DateOfBirth: user.DateOfBirth,
	}
	models.ManageResponse(w, "User created successfully.", http.StatusAccepted, responseModel, true)
}
