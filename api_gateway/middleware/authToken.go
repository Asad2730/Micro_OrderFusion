package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Asad2730/Micro_OrderFusion/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		token := tokenParts[1]
		if !isValidProductToken(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Optionally, you can extract user information from the token here
		claims, _ := getClaimsFromToken(token)
		c.Set("userID", claims["id"])
		c.Set("email", claims["email"])

		// Allow the request to proceed
		c.Next()
	}
}

func isValidProductToken(tokenString string) bool {
	jwtSecret := common.GetJWTSecret()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return false
	}

	return true
}

func getClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	jwtSecret := common.GetJWTSecret()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
