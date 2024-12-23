package routes

import (
	"real-time-chat-app/controllers"
	"real-time-chat-app/security"

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
		auth.POST("/login", func(c *gin.Context) {

			// Call SignUpController with ResponseWriter and Request
			controllers.LoginController(c.Writer, c.Request)
		})
		auth.Use(security.GinAuthMiddleware())
		{
			auth.POST("/logout", func(c *gin.Context) {

				// Call SignUpController with ResponseWriter and Request
				controllers.LogoutController(c)
			})
		}
	}
}

// SecureRoutes registers the secure routes with middleware applied
func SecureRoutes(r *gin.Engine) {
	// Define the /secure group
	secure := r.Group("/secure")
	secure.Use(security.GinAuthMiddleware()) // Apply authentication middleware to secure routes
	{
		// Define a secure route that returns secure data
		secure.GET("/data", controllers.SecureEndpoint)
	}
}
