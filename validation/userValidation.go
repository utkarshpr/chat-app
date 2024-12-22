package validation

import (
	"errors"
	"real-time-chat-app/models"
	"regexp"
	"unicode"
)

func SignUpUserValidation(user *models.User) error {

	// Validate username (at least 3 characters long)
	if len(user.Username) < 5 {
		return errors.New("username must be at least 5 characters long")
	}

	// Validate email format using a regular expression
	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	// Validate password (at least 6 characters long and contains at least one number)
	if len(user.Password) < 6 || !hasNumber(user.Password) || !hasSpecialCharacter(user.Password) {
		return errors.New("password must be at least 6 characters long, contain at least one number, and one special character")
	}

	// Validate first name (required)
	if user.FirstName == "" {
		return errors.New("first name is required")
	}

	// Validate last name (required)
	if user.LastName == "" {
		return errors.New("last name is required")
	}

	// Validate date of birth (required and must be in the format YYYY-MM-DD)
	if user.DateOfBirth == "" {
		return errors.New("date of birth is required")
	}

	// Validate address (optional but if provided, must not be empty)
	if user.Address != "" && len(user.Address) < 5 {
		return errors.New("address should be at least 5 characters long")
	}

	// If all validations pass, return nil
	return nil
}

func LoginUserValidation(user *models.LoginUser) error {

	// Validate username (at least 3 characters long)
	if len(user.Username) < 1 {
		return errors.New("username must be provided")
	}

	// Validate password (at least 6 characters long and contains at least one number)
	if len(user.Password) < 1 {
		return errors.New("password must be provided")
	}
	return nil
}

// Helper function to check if the email is valid
func isValidEmail(email string) bool {
	// Simple email regex pattern
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

// Helper function to check if the password contains at least one number
func hasNumber(s string) bool {
	for _, c := range s {
		if unicode.IsDigit(c) {
			return true
		}
	}
	return false
}

// Helper function to check if the password contains at least one special character
func hasSpecialCharacter(s string) bool {
	// Special characters include characters that are not alphanumeric
	re := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
	return re.MatchString(s)
}
