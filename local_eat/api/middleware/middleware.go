package middleware

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"local_eat/api/db/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func DBMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("db", db)
		context.Next()
	}
}

func AuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString, err := context.Cookie("token")
		if err != nil {
			context.AbortWithStatus(http.StatusUnauthorized)
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil {
			context.AbortWithStatus(http.StatusUnauthorized)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				context.AbortWithStatus(http.StatusUnauthorized)
			}
			user, err := auth.GetUser(db, claims["username"].(string))
			if err != nil {
				context.AbortWithStatus(http.StatusUnauthorized)
			}
			if user.Username == claims["username"] {
				context.Set("user", user)
				context.Next()
			}
		} else {
			context.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}