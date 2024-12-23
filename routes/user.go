package routes

import (
	"real-time-chat-app/controllers"
	"real-time-chat-app/logger"
	"real-time-chat-app/security"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	logger.LogInfo("User Routes ...")
	user := r.Group("/user")
	{
		user.Use(security.GinAuthMiddleware())
		{
			user.GET("/fetchUser", func(c *gin.Context) {

				// Call SignUpController with ResponseWriter and Request
				controllers.FetchUserController(c)
			})
		}
	}
}
