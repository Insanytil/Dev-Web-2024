package users

import (
	"net/http"
	"strconv"

	"local_eat/api/initializers"
	"local_eat/api/middleware"
	"local_eat/api/models"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	users := route.Group("/api/users")
	{
		users.GET("", middleware.AuthMiddleware, GetUsers)
		users.GET("get-company", middleware.AuthMiddleware, GetCompany)
		users.GET("get-producer", middleware.AuthMiddleware, GetProducer)
		users.POST("/create-company", middleware.AuthMiddleware, CreateCompany)
	}
}

func GetUsers(context *gin.Context) {
	user, ok := context.Get("user")
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}
	username := *user.(models.Users).Username

	var foundUser models.Users
	result := initializers.DB.Where("username = ?", username).First(&foundUser)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user from database"})
		return
	}
	if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	context.JSON(http.StatusOK, foundUser)
}
func GetProducer(context *gin.Context) {
	user, ok := context.Get("user")
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}
	foundUser := user.(models.Users).Username
	var producer models.Producers
	result := initializers.DB.Where("username = ?", *foundUser).First(&producer)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving producer from database"})
		return
	}
	if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Producer not found"})
		return
	}
	context.JSON(http.StatusOK, producer)
}
func GetCompany(context *gin.Context) {
	user, ok := context.Get("user")
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}
	foundUser := user.(models.Users).Username
	var producer models.Producers
	result := initializers.DB.Where("username = ?", *foundUser).First(&producer)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving producer from database"})
		return
	}
	if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Producer not found"})
		return
	}
	var relCompProd models.RelCompProd
	result2 := initializers.DB.Where("producer_id = ?", producer.ID).First(&relCompProd)
	if result2.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving CompanyName from database"})
		return
	}
	var company models.Company
	result3 := initializers.DB.Where("company_name = ?", relCompProd.CompanyName).First(&company)
	if result3.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving company from database"})
		return
	}
	context.JSON(http.StatusOK, company)
}

func CreateCompany(context *gin.Context) {
	user, ok := context.Get("user")
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}
	foundUser := user.(models.Users).Username

	var body models.Company
	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var producer models.Producers
	result := initializers.DB.Where("username = ?", *foundUser).First(&producer)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving producer from database"})
		return
	}
	if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Producer not found"})
		return
	}

	newCompany := models.Company{
		CompanyName: body.CompanyName,
		Password:    body.Password,
		Alias:       body.Alias,
		Address:     body.Address,
		Mail:        body.Mail,
		PhoneNum:    body.PhoneNum,
		VATNum:      body.VATNum,
		Description: body.Description,
	}
	companyResult := initializers.DB.Create(&newCompany)
	if companyResult.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	newRel := models.RelCompProd{
		ProducerID:  strconv.Itoa(producer.ID),
		CompanyName: body.CompanyName,
	}
	relResult := initializers.DB.Create(&newRel)
	if relResult.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{})
}
