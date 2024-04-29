package users

import (
	"net/http"

	"local_eat/api/initializers"
	model "local_eat/api/models"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	users := route.Group("/api/producers")
	{
		users.GET("", GetProducers)
		users.POST("/register", RegisterProducers)
	}
}

// swagger:operation GET /api/producers Producers GetProducersRequest
// GET Producers
// @Summary Get producers
// @Description Get producers id, name, picture and created values
// @Tags Producers
// @Produce json
// @Success 200 {array} model.Producers
// @Failure 404 "Not found"
// @Failure 500 "Internal server error"
// @Router /api/producers [get]
func GetProducers(context *gin.Context) {
	var producers []*model.Producers
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
// @Param body body model.Producers true "Producer object to be registered"
// @Success 200
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /api/producers/register [post]
func RegisterProducers(context *gin.Context) {
	var body model.Producers
	if context.BindJSON(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	userInterface, exists := context.Get("user")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found in context",
		})
		return
	}
	existingUser, ok := userInterface.(model.Users)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to convert user from context",
		})
		return
	}

	// Créez une nouvelle instance de Producers avec les données envoyées depuis le frontend
	producers := model.Producers{
		Users:     existingUser,
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		PhoneNum:  body.PhoneNum,
		EmailPro:  body.EmailPro,
	}

	// Enregistrez le producteur dans la base de données
	result := initializers.DB.Create(&producers)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	// Répondez avec un statut OK si tout s'est bien passé
	context.JSON(http.StatusOK, gin.H{})
}
