package security

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// GinAuthMiddleware is a middleware that checks if the JWT token is valid
func GinAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user", claims)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to parse claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// SecureEndpoint handles the secure endpoint requests
func SecureEndpoint(c *gin.Context) {
	// Retrieve the user data from the context
	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Return secure data along with user claims
	c.JSON(http.StatusOK, gin.H{
		"message": "This is a secure endpoint.",
		"user":    userClaims,
	})
}
