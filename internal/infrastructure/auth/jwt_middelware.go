package auth

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(jwtService JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not found"})
			c.Abort()
			return
		}

		// Split the Authorization header to extract the token
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		// Validate the token
		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check if token is valid and has claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// You can extract user information from the token and set it in the context
			userID := claims["user_id"].(string)
			c.Set("user_id", userID)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next() // Proceed to the next middleware or the handler
	}
}
