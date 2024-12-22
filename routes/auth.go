package routes

import (
	"real-time-chat-app/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		// Wrapping SignUpController to work with Gin's gin.Context
		auth.POST("/signup", func(c *gin.Context) {

			// Call SignUpController with ResponseWriter and Request
			controllers.SignUpController(c.Writer, c.Request)
		})
		// Uncomment and modify the following if needed
		// auth.POST("/login", controllers.Login)
		// auth.POST("/logout", controllers.Logout)
	}
}
