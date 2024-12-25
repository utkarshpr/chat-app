package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"real-time-chat-app/models"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var (
	Connections = make(map[string]*websocket.Conn) // Map of recipientID to WebSocket connection
	ConnMutex   sync.Mutex                         // Mutex for concurrent access to connections map
)

// GenerateUUID generates a unique identifier
func GenerateUUID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes)
}

// GetCurrentTimestamp returns the current timestamp in ISO 8601 format
func GetCurrentTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// EstablishWebSocketConnection sets up a WebSocket connection
func EstablishWebSocketConnection(w http.ResponseWriter, r *http.Request) {
	recipientID := r.URL.Query().Get("recipient_id")
	if recipientID == "" {
		log.Println("Recipient ID is required for WebSocket connection")
		http.Error(w, "Recipient ID is required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade WebSocket:", err)
		return
	}
	defer UnregisterConnection(recipientID)

	RegisterConnection(recipientID, conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading WebSocket message:", err)
			break
		}
		log.Printf("Received message from %s: %s", recipientID, message)
	}
}

// BroadcastToRecipient sends a message to the recipient via WebSocket
func BroadcastToRecipient(recipientID string, message *models.Message) {
	ConnMutex.Lock()
	defer ConnMutex.Unlock()

	conn, exists := Connections[recipientID]
	if !exists {
		log.Printf("Recipient %s is not online. Message cannot be delivered in real-time.", recipientID)
		return
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to serialize message for recipient %s: %v", recipientID, err)
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, messageBytes)
	if err != nil {
		log.Printf("Failed to send message to recipient %s: %v", recipientID, err)
		delete(Connections, recipientID)
		conn.Close()
	}
}

// RegisterConnection registers a WebSocket connection for a user
func RegisterConnection(recipientID string, conn *websocket.Conn) {
	ConnMutex.Lock()
	defer ConnMutex.Unlock()

	if existingConn, exists := Connections[recipientID]; exists {
		existingConn.Close()
	}

	Connections[recipientID] = conn
	log.Printf("WebSocket connection established for recipient %s", recipientID)
}

// UnregisterConnection unregisters a WebSocket connection for a user
func UnregisterConnection(recipientID string) {
	ConnMutex.Lock()
	defer ConnMutex.Unlock()

	if conn, exists := Connections[recipientID]; exists {
		conn.Close()
		delete(Connections, recipientID)
		log.Printf("WebSocket connection closed for recipient %s", recipientID)
	}
}

// HandleError formats and sends an error response
func HandleError(c *gin.Context, statusCode int, message string, err error) {
	log.Printf("Error: %s, Details: %v", message, err)
	c.JSON(statusCode, gin.H{"error": message, "details": err.Error()})
}
