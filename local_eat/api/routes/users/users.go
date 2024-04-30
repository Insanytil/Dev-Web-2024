package users

import (
	"net/http"

	"local_eat/api/initializers"
	"local_eat/api/middleware"
	"local_eat/api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(route *gin.Engine) {
	users := route.Group("/api/producers")
	{
		users.GET("", GetProducers)
		users.POST("/register", middleware.AuthMiddleware, RegisterProducers)
	}
}

// swagger:operation GET /api/producers Producers GetProducersRequest
// GET Producers
// @Summary Get producers
// @Description Get producers id, name, picture and created values
// @Tags Producers
// @Produce json
// @Success 200 {array} models.Producers
// @Failure 404 "Not found"
// @Failure 500 "Internal server error"
// @Router /api/producers [get]
func GetProducers(context *gin.Context) {
	var producers []*models.Producers
	result := initializers.DB.Find(&producers)
	if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{})
		return

	}
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusOK, producers)
}

// swagger:operation POST /api/producers/register PostProducersRequest
// POST producers
// @Summary POST producers
// @Description Post producer Lastname, Firstname, Phone number and pro email
// @Tags Producers
// @Accept json
// @Produce json
// @Param body body models.Producers true "Producer object to be registered"
// @Success 201 "Producteur created"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /api/producers/register [post]
func RegisterProducers(context *gin.Context) {
	// Récupérer le nom d'utilisateur du contexte
	user, _ := context.Get("user")

	var body models.Producers
	if context.BindJSON(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	newProducer := models.Producers{
		Username:  *user.(models.Users).Username,
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		PhoneNum:  body.PhoneNum,
		EmailPro:  body.EmailPro,
	}

	var oldProducer models.Producers
	exists := initializers.DB.First(&oldProducer, "username = ?", newProducer.Username)
	if exists.Error != gorm.ErrRecordNotFound {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "User already a producer",
		})
		return
	}

	result := initializers.DB.Create(&newProducer)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{})
}
