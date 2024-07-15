package middleware

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func AuthMiddleware(jwtSecretKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString := ctx.Request.Header.Get("Authorization")
		log.Println(tokenString)
		if tokenString == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Missing or invalid token"})
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})

		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid token: " + err.Error()})
			return
		}

		if !token.Valid {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		ctx.Set("claims", claims)

		ctx.Next()
	}
}
