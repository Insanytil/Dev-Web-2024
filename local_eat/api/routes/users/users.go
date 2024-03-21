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
