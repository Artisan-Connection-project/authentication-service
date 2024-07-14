package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtSecretKey = []byte("secret-key")

type claims struct {
	jwt.StandardClaims
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Missing or invalid token"})
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		ctx.Next()
	}
}
