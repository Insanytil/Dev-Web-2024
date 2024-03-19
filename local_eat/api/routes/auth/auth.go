package auth

import (
	"net/http"
	"os"
	"time"

	"local_eat/api/initializers"
	"local_eat/api/middleware"
	"local_eat/api/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Routes(route *gin.Engine) {
	users := route.Group("/api/auth")
	{
		users.POST("/signup", signup)
		users.POST("/login", login)
		users.GET("/authenticate", middleware.AuthMiddleware, Authenticate)
	}
}

// swagger:operation POST /api/auth/signup Auth PostSignupRequest
// POST Signup
// @Summary Send user data to create a new user
// @Description Send user data to create a new user
// @Tags Auth
// @Accept json
// @Param user body model.Users true "User data"
// @Success 200 "User created"
// @Failure 400 "Invalid request"
// @Failure 500 "Internal server error"
// @Router /api/auth/signup [post]
func signup(context *gin.Context) {
	var body model.Users
	if context.BindJSON(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error"})
		return
	}

	user := model.Users{
		Username: body.Username,
		Password: string(hash),
		Email:    body.Email,
	}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error"})
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
// @Param user body model.Users true "User data"
// @Success 200 "User authenticated"
// @Failure 400 "User not found"
// @Failure 400 "Invalid password"
// @Failure 500 "Internal server error"
// @Router /api/auth/login [post]
func login(context *gin.Context) {
	var body model.Users
	if context.BindJSON(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request"})
		return
	}

	var user model.Users
	var result *gorm.DB
	if body.Username == nil {
		result = initializers.DB.First(&user, "email = ?", body.Email)
	} else {
		result = initializers.DB.First(&user, "username = ?", body.Username)
	}

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
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
// @Success 200 "User authenticated"
// @Failure 401 "Unauthorized"
// @Router /api/auth/authenticate [get]
func Authenticate(context *gin.Context) {
	user, _ := context.Get("user")
	context.JSON(http.StatusOK, gin.H{
		"user": user.(model.Users).Username,
	})
}