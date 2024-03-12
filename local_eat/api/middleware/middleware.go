package middleware

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func DBMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("db", db)
		context.Next()
	}
}
