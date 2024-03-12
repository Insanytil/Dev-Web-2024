package middleware

import (
	"github.com/gin-gonic/gin"
	"local_eat/api/db"
)

func DBMiddleware(db db.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("db", db)
		context.Next()
	}
}
