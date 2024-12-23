package main

import (
	"log"
	"os"
	"real-time-chat-app/database"
	"real-time-chat-app/logger"
	repo "real-time-chat-app/repositary"

	routes "real-time-chat-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	logger.InitLogger("app.log")

	loadEnvVarible()
	// Initialize MongoDB connection
	database.InitMongoDB()
	repo.InitRepository()

	// Set up the Gin router
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	// Register authentication routes
	routes.AuthRoutes(r)
	routes.SecureRoutes(r)
	routes.UserRoutes(r)

	port := os.Getenv("PORT")
	err := r.Run(port) // Changes the port to 8081
	if err != nil {
		logger.LogError("Failed to start server: " + err.Error())
	}

	//TestDatabase()

}

func loadEnvVarible() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
