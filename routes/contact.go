package routes

import (
	"real-time-chat-app/controllers"
	"real-time-chat-app/security"

	"github.com/gin-gonic/gin"
)

func ContactRoutes(r *gin.Engine) {
	contact := r.Group("/contact")
	{
		contact.Use(security.GinAuthMiddleware())
		{
			contact.POST("/add", func(c *gin.Context) {

				// Call SignUpController with ResponseWriter and Request
				controllers.AddAndUpdateCOntact(c)
			})

			contact.GET("/get", func(c *gin.Context) {

				// Call SignUpController with ResponseWriter and Request
				controllers.GetListofContact(c)

			})

			contact.POST("/action", func(c *gin.Context) {
				controllers.BlockOrRemoveContact(c)
			})
		}

	}

}
