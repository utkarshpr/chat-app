package logger

import (
	"log"
	"os"
)

// Logger is the global logger object
var Logger *log.Logger

// InitLogger initializes the logger with a file output
func InitLogger(logFile string) {
	// Open the log file, create it if it doesn't exist, append to it if it does
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	// Initialize the logger to write to the file with timestamp prefix
	Logger = log.New(file, "", log.LstdFlags)
	// Log that the logger has been initialized
	Logger.Println("Logger initialized.")
}

// LogInfo logs an info-level message
func LogInfo(message string) {
	if Logger != nil {
		Logger.Println("INFO:", message)
	} else {
		log.Println("Logger not initialized!")
	}
}

// LogError logs an error-level message
func LogError(message string) {
	if Logger != nil {
		Logger.Println("ERROR:", message)
	} else {
		log.Println("Logger not initialized!")
	}
}
