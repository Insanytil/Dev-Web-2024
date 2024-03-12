package auth

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"local_eat/api/db/auth"
	"local_eat/api/middleware"
	"local_eat/api/model"
	"net/http"
)

func Routes(route *gin.Engine, db *sql.DB) {
	users := route.Group("/api/auth")
	{
		users.Use(middleware.DBMiddleware(db))
		users.POST("/signup", signup)
	}
}

// swagger:operation POST /api/auth/signup Auth PostSignupRequest
// POST Signup
// @Summary Send user data to create a new user
// @Description Send user data to create a new user
// @Tags Auth
// @Accept json
// @Param user body model.Users true "User data"
// @Success 200 {string} string ""
// @Failure 500 {string} string "Internal server error"
// @Router /api/auth/signup [post]
func signup(context *gin.Context) {
	db := context.MustGet("db").(*sql.DB)
	var user model.Users
	if context.BindJSON(&user) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error"})
		return
	}

	user.Password = string(hash)
	err = auth.CreateUser(db, &user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{})
}
