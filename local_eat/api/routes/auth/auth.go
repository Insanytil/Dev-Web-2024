package auth

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"local_eat/api/db/auth"
	"local_eat/api/middleware"
	"local_eat/api/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Routes(route *gin.Engine, db *sql.DB) {
	users := route.Group("/api/auth")
	{
		users.Use(middleware.DBMiddleware(db))
		users.POST("/signup", signup)
		users.POST("/login", login)
		users.GET("/authenticate", middleware.AuthMiddleware(db), Authenticate)
	}
}

// swagger:operation POST /api/auth/signup Auth PostSignupRequest
// POST Signup
// @Summary Send user data to create a new user
// @Description Send user data to create a new user
// @Tags Auth
// @Accept json
// @Param user body model.UsersSignup true "User data"
// @Success 200 {string} string ""
// @Failure 400 {string} string "Invalid request"
// @Failure 400 {string} string "Failed to create user"
// @Failure 500 {string} string "Internal server error"
// @Router /api/auth/signup [post]
func signup(context *gin.Context) {
	db := context.MustGet("db").(*sql.DB)
	var user model.UsersSignup
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

// swagger:operation POST /api/auth/login Auth PostLoginRequest
// POST Login
// @Summary Send username and password to login
// @Description Send username and password to login to receive a token in a cookie
// @Tags Auth
// @Accept json
// @Param user body model.UsersLogin true "User data"
// @Success 200 {string} string ""
// @Failure 400 {string} string "User not found"
// @Failure 400 {string} string "Invalid password"
// @Failure 500 {string} string "Internal server error"
// @Router /api/auth/login [post]
func login(context *gin.Context) {
	db := context.MustGet("db").(*sql.DB)
	var user model.UsersLogin
	if context.BindJSON(&user) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request"})
		return
	}

	userDB, err := auth.GetUser(db, user.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error"})
		return
	}
	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("token", tokenString, 3600 * 24, "", "", false, true)
	context.JSON(http.StatusOK, gin.H{})
}

// swagger:operation GET /api/auth/authenticate Auth GetAuthenticateRequest
// GET Authenticate
// @Summary Validate user token
// @Description Validate user token
// @Tags Auth
// @Produce json
// @Success 200 {string} string "User authenticated"
// @Failure 401 {string} string "Unauthorized"
// @Router /api/auth/authenticate [get]
func Authenticate(context *gin.Context) {
	user, _ := context.Get("user")
	context.JSON(http.StatusOK, gin.H{
		"user": user.(model.UsersLogin).Username,
	})
}