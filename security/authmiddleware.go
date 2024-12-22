package security

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// GinAuthMiddleware wraps AuthMiddleware for Gin framework
func GinAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

		// Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Extract claims and store in context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user", claims)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to parse claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// SecureEndpoint handler for secure routes
func SecureEndpoint(c *gin.Context) {
	// Fetch user data from context
	userClaims, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Return secure data
	c.JSON(http.StatusOK, gin.H{
		"message": "This is a secure endpoint.",
		"user":    userClaims,
	})
}
