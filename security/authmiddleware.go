package security

import (
	"fmt"
	"net/http"
	"os"
	"real-time-chat-app/logger"
	"real-time-chat-app/models"
	repo "real-time-chat-app/repositary"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// GinAuthMiddleware is a middleware that checks if the JWT token is valid
func GinAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.LogInfo("GinAuthMiddleware ... :: started")
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
		// Parse and validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logger.LogError("unexpected signing method")
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract claims and store them in context
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			c.Set("user", claims)
		} else {
			models.ManageResponse(c.Writer, "Unable to parse claims", http.StatusBadRequest, nil, false)
			c.Abort()
			return
		}

		username, ok := claims["username"].(string)
		if !ok || username == "" {
			models.ManageResponse(c.Writer, "Invalid claims: username missing", http.StatusBadRequest, nil, false)
			c.Abort()
			return
		}

		// Check if the session exists in jwtCollection
		_, err = repo.FetchJwtTokenForUser(username)
		if err != nil {
			logger.LogError("Session not found for user: " + username)
			models.ManageResponse(c.Writer, "Session expired or Login again", http.StatusNonAuthoritativeInfo, nil, false)
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
		logger.LogInfo("GinAuthMiddleware ... :: ended")
	}
}

func GetClaims(c *gin.Context) jwt.MapClaims {
	authHeader := c.GetHeader("Authorization")
	logger.LogInfo("GetClaims :: started")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		c.Abort()
		return nil
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.LogInfo("unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return nil
	}

	// Extract claims and store them in context
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		models.ManageResponse(c.Writer, "Unable to parse claims", http.StatusBadRequest, nil, false)
		c.Abort()
		return nil
	}
	logger.LogInfo("GetClaims :: ended")
	return claims

}
