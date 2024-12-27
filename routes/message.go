package routes

import (
	"log"
	"net/http"
	"real-time-chat-app/controllers"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	"real-time-chat-app/security"
	"real-time-chat-app/utils"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)

func MessageRoute(r *gin.Engine) {
	logger.LogInfo("MessageRoute Routes ...")
	user := r.Group("/message")
	{
		user.Use(security.GinAuthMiddleware())
		{
			user.POST("/sent", func(c *gin.Context) {

				// Call SignUpController with ResponseWriter and Request
				controllers.MessageSentController(c)
			})
			user.PATCH("/edit", func(c *gin.Context) {

				// Call SignUpController with ResponseWriter and Request
				controllers.MessageEditController(c)
			})

			user.DELETE("/delete", func(c *gin.Context) {

				// Call SignUpController with ResponseWriter and Request
				controllers.MessageDeleteController(c)
			})

			user.GET("/get", func(c *gin.Context) {

				// Call SignUpController with ResponseWriter and Request
				controllers.MessageGetAllController(c)
			})

		}

	}

}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// This allows connections from any origin. You can customize this for security.
		return true
	},
}

func WebSocketRoute(r *gin.Engine) {
	r.GET("/ws", func(c *gin.Context) {
		log.Println("WebSocket connection requested")

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Failed to upgrade connection:", err)
			return
		}

		userID := c.DefaultQuery("userID", "")
		if userID == "" {
			log.Println("No userID provided")
			models.ManageResponse(c.Writer, "No userID provided", http.StatusBadRequest, nil, false)
			conn.Close()
			return
		}

		utils.ConnMutex.Lock()
		utils.Connections[userID] = conn
		utils.ConnMutex.Unlock()

		log.Printf("WebSocket connection established for %s", userID)

		// Handle closure
		defer func() {
			utils.ConnMutex.Lock()
			delete(utils.Connections, userID)
			utils.ConnMutex.Unlock()
			conn.Close()
			log.Printf("WebSocket connection closed for %s", userID)
		}()

		// Listen for messages and handle them
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Error reading message: %v", err)
				break
			}

			// Optionally send a message back
			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Printf("Error sending message: %v", err)
				break
			}
		}
	})

}
