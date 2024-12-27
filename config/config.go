package config

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"mime/multipart"
	"os"
	"real-time-chat-app/logger"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func GenerateStrongSecretKey(length int) string {
	// Create a random byte slice of the given length
	secret := make([]byte, length)
	_, err := rand.Read(secret)
	if err != nil {
		log.Fatal("Error generating random bytes:", err)
	}

	// Return the base64 encoded string of the random bytes
	return base64.StdEncoding.EncodeToString(secret)
}

func InitCloudinary() (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
		return nil, err
	}
	logger.LogInfo("InitCloudinary successfull ")
	return cld, nil
}

func UploadMedia(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	// Initialize Cloudinary
	cld, err := InitCloudinary()
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	// Upload the file to Cloudinary
	uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: "message_media",
	})
	if err != nil {
		log.Printf("Failed to upload media to Cloudinary: %v", err)
		return "", err
	}

	// Return the media URL
	return uploadResult.SecureURL, nil
}
