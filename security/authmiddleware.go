package security

import (
	"fmt"
	"net/http"
	"os"
	"real-time-chat-app/logger"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// GinAuthMiddleware is a middleware that checks if the JWT token is valid
func GinAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		logger.LogInfo(authHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
		logger.LogInfo(tokenString)
		// Parse and validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logger.LogInfo("unexpected signing method")
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secretKey, nil
		})
		logger.LogInfo("...........................")
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract claims and store them in context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user", claims)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to parse claims"})
			c.Abort()
			return
		}

		// Safely extract the role
		// userRoleStr, exists := claims["role"].(string)
		// if !exists {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "role not found in token"})
		// 	c.Abort()
		// 	return
		// }

		// // Parse the role string into the Role type
		// userRole, err := models.
		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid role"})
		// 	c.Abort()
		// 	return
		// }
		// Store the role and claims in the context

	}
}
