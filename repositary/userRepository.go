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
	"golang.org/x/crypto/bcrypt"
)

// Declare the collection globally, but initialize it after InitMongoDB is called
var userCollection *mongo.Collection

// Initialize userCollection after the MongoDB connection is established
func InitRepository() {
	database.InitMongoDB()

	userCollection = database.GetCollection(os.Getenv("MONGO_TABLE_USER"))
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

func IsLoggedinUserExist(user *models.LoginUser) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find user by email
	var existingUser models.User
	filter := bson.M{"username": user.Username}
	err := userCollection.FindOne(ctx, filter).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("user not found")
		}
		return "", err
	}

	logger.LogInfo("IsLoggedinUserExist :: user fetch from username : " + existingUser.Email)

	// Check if the provided password matches the stored hashed password
	err = CompareHashAndPassword(existingUser.Password, user.Password)

	if err != nil {
		return "", errors.New("invalid password")
	}

	logger.LogInfo("IsLoggedinUserExist :: Hashed password check success : ")

	// Generate a JWT token here (you can use any JWT library to generate the token)
	// You can use `jwt-go` or `golang-jwt/jwt` for this purpose.
	token, err := generateJWT(existingUser)
	if err != nil {
		return "", err
	}

	return token, nil
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
		logger.LogInfo("generateJWT :: error in  Create a new token with claims and sign it with the secret key" + err.Error())
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
