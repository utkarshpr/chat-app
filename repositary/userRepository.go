package repo

import (
	"context"
	"errors"
	"os"
	"real-time-chat-app/database"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// Declare the collection globally, but initialize it after InitMongoDB is called
var userCollection *mongo.Collection
var jwtCollection *mongo.Collection
var contactCollection *mongo.Collection

// Initialize userCollection after the MongoDB connection is established
func InitRepository() {
	database.InitMongoDB()

	userCollection = database.GetCollection(os.Getenv("MONGO_TABLE_USER"))
	jwtCollection = database.GetCollection(os.Getenv("MONGO_TABLE_JWT_STORE"))
	contactCollection = database.GetCollection(os.Getenv("MONGO_TABLE_CONTACT"))
}

// InsertUser inserts a new user into the database
func InsertUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"$or": []bson.M{
			{"email": user.Email},
			{"username": user.Username},
		},
	}

	count, err := userCollection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("email or username already exists")
	}

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func IsLoggedinUserExist(user *models.LoginUser) (string, string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find user by email
	var existingUser models.User
	filter := bson.M{"username": user.Username}
	err := userCollection.FindOne(ctx, filter).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", "", errors.New("user not found")
		}
		return "", "", err
	}

	logger.LogInfo("IsLoggedinUserExist :: user fetch from username : " + existingUser.Email)

	// Check if the provided password matches the stored hashed password
	err = CompareHashAndPassword(existingUser.Password, user.Password)

	if err != nil {
		return "", "", errors.New("invalid password")
	}

	logger.LogInfo("IsLoggedinUserExist :: Hashed password check success : ")

	// Generate a JWT token here (you can use any JWT library to generate the token)
	// You can use `jwt-go` or `golang-jwt/jwt` for this purpose.
	token, err := generateJWT(existingUser)
	if err != nil {
		return "", "", err
	}
	refreshtoken, err := generateRefreshTokenJWT(existingUser)
	if err != nil {
		return "", "", err
	}

	tok, ref, err := InsertJwtTokenForUser(token, refreshtoken, &existingUser)
	if err != nil {
		logger.LogError("IsLoggedinUserExist :: Unable to store JWT token in DB ")
		return "", "", errors.New("unable to store jwt token in db ")
	}
	return tok, ref, nil
}

func generateJWT(user models.User) (string, error) {
	// Create claims with user data
	claims := jwt.MapClaims{
		"user_id":    user.ID.Hex(), // User ID (in case you want to identify the user by ID)
		"username":   user.Username, // User's username
		"First Name": user.FirstName,
		"Last Name":  user.LastName,
		"role":       user.Role,                             // User's email
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (1 day)
	}

	logger.LogInfo("generateJWT :: claim map formed ")

	// Create a new token with claims and sign it with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		logger.LogError("generateJWT :: error in  Create a new token with claims and sign it with the secret key" + err.Error())
		return "", err
	}

	return tokenString, nil
}

func generateRefreshTokenJWT(user models.User) (string, error) {
	// Create claims with user data
	claims := jwt.MapClaims{
		"user_id":    user.ID.Hex(), // User ID (in case you want to identify the user by ID)
		"username":   user.Username, // User's username
		"First Name": user.FirstName,
		"Last Name":  user.LastName,
		"role":       user.Role,                                 // User's email
		"exp":        time.Now().Add(time.Hour * 24 * 7).Unix(), // Token expiration time (7 day)
	}

	logger.LogInfo("generateRefreshTokenJWT :: claim map formed ")

	// Create a new token with claims and sign it with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		logger.LogError("generateRefreshTokenJWT :: error in  Create a new token with claims and sign it with the secret key" + err.Error())
		return "", err
	}

	return tokenString, nil
}

// FetchUserByUsername fetches a user from the database using their username.
func FetchUserByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": username}

	var user models.User
	err := database.GetCollection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func CompareHashAndPassword(fetchedUserPassword string, loginUserPassword string) error {
	// Compare the stored hash with the provided password
	err := bcrypt.CompareHashAndPassword([]byte(fetchedUserPassword), []byte(loginUserPassword))
	if err != nil {
		return errors.New("invalid credentials") // Password does not match
	}
	return nil
}

func InsertJwtTokenForUser(token string, refreshToken string, user *models.User) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": user.Username}

	loginResp := &models.LoginResponse{
		Username:     user.Username,
		RefreshToken: refreshToken,
		Token:        token,
	}
	update := bson.M{
		"$set": loginResp, // Replace the document with the new data
	}
	opts := options.Update().SetUpsert(true)
	_, err := jwtCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		logger.LogError("InsertJwtTokenForUser :: Unable to store JWT token in DB " + err.Error())
		return "", "", err
	}
	return token, refreshToken, nil
}

func FetchJwtTokenForUser(username string) (*models.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the filter to find the user's token
	filter := bson.M{"username": username}

	// Define a variable to hold the result
	var loginResp models.LoginResponse

	// Perform the query to find the document
	err := jwtCollection.FindOne(ctx, filter).Decode(&loginResp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.LogError("FetchJwtTokenForUser :: No JWT token found for user: " + username)
			return nil, errors.New("no JWT token found for the user")
		}
		logger.LogError("FetchJwtTokenForUser :: Error fetching JWT token: " + err.Error())
		return nil, err
	}

	logger.LogInfo("FetchJwtTokenForUser :: Successfully fetched JWT token for user: " + username)
	return &loginResp, nil
}

func LogoutUser(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the filter to identify the user's token
	filter := bson.M{"username": username}

	// Attempt to delete the user's token document from the collection
	result, err := jwtCollection.DeleteOne(ctx, filter)
	if err != nil {
		logger.LogError("Logout :: Error while removing JWT token from DB for user: " + username + " - " + err.Error())
		return errors.New("failed to logout the user")
	}

	// Check if any document was deleted
	if result.DeletedCount == 0 {
		logger.LogError("Logout :: No JWT token found to remove for user: " + username)
		return errors.New("no active session found for the user")
	}

	logger.LogInfo("Logout :: Successfully logged out user: " + username)
	return nil
}

func UserFetchFromDB(username string) (*models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.UserResponse
	filter := bson.M{"username": username}

	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func UserAndProfileUpdate(username string, updateUser *models.UpdateUserAndProfile) (*models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	filter := bson.M{"username": username}

	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.LogError("UserAndProfileUpdate :: user not found")
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	err = handleInvalidDatainUpdate(updateUser, &user)
	updateData := bson.M{
		"password":       updateUser.Password,
		"avatar_url":     updateUser.AvatarURL,
		"status_message": updateUser.StatusMessage,
		"last_seen":      updateUser.LastSeen,
		"first_name":     updateUser.FirstName,
		"last_name":      updateUser.LastName,
		"address":        updateUser.Address,
		"date_of_birth":  updateUser.DateOfBirth,
		"role":           updateUser.Role,
		"profile": bson.M{
			"bio":                  updateUser.Profile.Bio,
			"is_profile_public":    updateUser.Profile.IsProfilePublic,
			"cover_photo_url":      updateUser.Profile.CoverPhotoURL,
			"profile_completeness": updateUser.Profile.ProfileCompleteness,
			"social_links":         updateUser.Profile.SocialLinks,
			"interests":            updateUser.Profile.Interests,
			"contact_preferences":  updateUser.Profile.ContactPreferences,
			"occupation":           updateUser.Profile.Occupation,
			"education":            updateUser.Profile.Education,
			"achievements":         updateUser.Profile.Achievements,
			"gender":               updateUser.Profile.Gender,
			"phone_number":         updateUser.Profile.PhoneNumber,
		},
	}

	update := bson.M{"$set": updateData}

	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var userResponse models.UserResponse
	// Fetch the updated user data
	err = userCollection.FindOne(ctx, filter).Decode(&userResponse)
	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}

func handleInvalidDatainUpdate(updateUser *models.UpdateUserAndProfile, user *models.User) error {
	// Overwrite fields in updateUser with existing user data if the field is empty
	if updateUser.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("failed to hash password")
		}
		updateUser.Password = string(hashedPassword)
	} else {
		// If password is empty, keep the existing password
		updateUser.Password = user.Password
	}

	if updateUser.FirstName == "" {
		updateUser.FirstName = user.FirstName
	}
	if updateUser.LastName == "" {
		updateUser.LastName = user.LastName
	}
	if updateUser.AvatarURL == "" {
		updateUser.AvatarURL = user.AvatarURL
	}
	if updateUser.StatusMessage == "" {
		updateUser.StatusMessage = user.StatusMessage
	}
	if updateUser.LastSeen == "" {
		updateUser.LastSeen = user.LastSeen
	}
	if updateUser.Address == "" {
		updateUser.Address = user.Address
	}
	if updateUser.DateOfBirth == "" {
		updateUser.DateOfBirth = user.DateOfBirth
	}
	if updateUser.Role == "" {
		updateUser.Role = user.Role
	}
	if updateUser.Profile.Bio == "" {
		updateUser.Profile.Bio = user.Profile.Bio
	}
	if updateUser.Profile.IsProfilePublic != true || updateUser.Profile.IsProfilePublic != false {
		updateUser.Profile.IsProfilePublic = user.Profile.IsProfilePublic
	}
	if updateUser.Profile.CoverPhotoURL == "" {
		updateUser.Profile.CoverPhotoURL = user.Profile.CoverPhotoURL
	}
	if updateUser.Profile.ProfileCompleteness == 0 {
		updateUser.Profile.ProfileCompleteness = user.Profile.ProfileCompleteness
	}
	if updateUser.Profile.SocialLinks == nil {
		updateUser.Profile.SocialLinks = user.Profile.SocialLinks
	}
	if updateUser.Profile.Interests == nil {
		updateUser.Profile.Interests = user.Profile.Interests
	}
	if updateUser.Profile.ContactPreferences == nil {
		updateUser.Profile.ContactPreferences = user.Profile.ContactPreferences
	}
	if updateUser.Profile.Occupation == "" {
		updateUser.Profile.Occupation = user.Profile.Occupation
	}
	if updateUser.Profile.Education == "" {
		updateUser.Profile.Education = user.Profile.Education
	}
	if updateUser.Profile.Achievements == nil {
		updateUser.Profile.Achievements = user.Profile.Achievements
	}
	if updateUser.Profile.Gender == "" {
		updateUser.Profile.Gender = user.Profile.Gender
	}
	if updateUser.Profile.PhoneNumber == "" {
		updateUser.Profile.PhoneNumber = user.Profile.PhoneNumber
	}
	return nil

}

func DeleteUser(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": username}

	deletedResult, err := userCollection.DeleteOne(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("user not found")
		}

		return nil
	}
	if deletedResult.DeletedCount == 0 {
		return errors.New("user not found with username " + username)
	}
	logger.LogInfo("User Deleted from the database" + string(deletedResult.DeletedCount))
	return nil
}
