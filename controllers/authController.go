package controllers

import (
	"encoding/json"
	"net/http"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	"real-time-chat-app/services"
	"real-time-chat-app/validation"
	"strings"

	"github.com/gin-gonic/gin"
)

// SignUpController handles user sign-up requests.
//
// @Summary Sign up a new user
// @Description Processes user registration by validating the input and creating a new user.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.User true "User Details"
// @Success 201 {object} models.Response{data=models.UserResponse} "User created successfully"
// @Failure 400 {object} models.Response "Invalid input or validation errors"
// @Failure 405 {object} models.Response "Method not allowed"
// @Failure 500 {object} models.Response "Internal server error"
// @Router /auth/signup [post]
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
		if strings.HasPrefix(err.Error(), "invalid role:") {
			models.ManageResponse(w, "Error : "+"Invalid role provided. Allowed roles are 'ADMIN' or 'CLIENT'.", http.StatusBadRequest, nil, false)
			return
		}
		logger.LogInfo("SignUpController :: error in decoding the body" + err.Error())
		models.ManageResponse(w, "error in decoding the body "+err.Error(), http.StatusBadRequest, nil, false)

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

// LoginController handles user login requests.
//
// @Summary User login
// @Description Authenticates the user with valid credentials and returns a JWT token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.LoginUser true "Login User Details"
// @Success 200 {object} models.Response{data=models.LoginResponse} "User logged in successfully"
// @Failure 400 {object} models.Response "Invalid input or validation errors"
// @Failure 405 {object} models.Response "Method not allowed"
// @Failure 500 {object} models.Response "Internal server error"
// @Router /auth/login [post]
func LoginController(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		logger.LogInfo("LoginController :: error POST method required")
		models.ManageResponse(w, "POST method required", http.StatusMethodNotAllowed, nil, false)
		return
	}

	var user models.LoginUser

	// Decode the JSON body into the User struct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		logger.LogInfo("LoginController :: error in decoding the body")
		models.ManageResponse(w, "error in decoding the body", http.StatusBadRequest, nil, false)

		return
	}

	// validation
	err = validation.LoginUserValidation(&user)
	if err != nil {
		logger.LogInfo("LoginController :: error in validation  " + err.Error())
		models.ManageResponse(w, "Error : "+err.Error(), http.StatusBadRequest, nil, false)
		return
	}

	//password match
	token, refreshtoken, err := services.LoginUser(&user)
	if err != nil {
		logger.LogInfo("LoginController :: error in service call " + err.Error())
		models.ManageResponse(w, "Unable to login the User "+err.Error(), http.StatusBadRequest, nil, false)
		return
	}

	resp := &models.LoginResponse{
		Username:     user.Username,
		Token:        token,
		RefreshToken: refreshtoken,
	}

	models.ManageResponse(w, "User LoggedIn successfully.", http.StatusOK, resp, true)

}

// SecureEndpoint handles the secure endpoint requests
func SecureEndpoint(c *gin.Context) {
	// Retrieve the user data from the context
	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Return secure data along with user claims
	c.JSON(http.StatusOK, gin.H{
		"message": "This is a secure endpoint.",
		"user":    userClaims,
	})
}
