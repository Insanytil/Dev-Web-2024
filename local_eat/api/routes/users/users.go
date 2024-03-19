package users

import (
	"net/http"

	"local_eat/api/initializers"
	"local_eat/api/model"

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
// @Failure 500 {string} string "Internal server error"
// @Router /api/producers [get]
func GetProducers(context *gin.Context) {
	var producers []*model.Producers
	initializers.DB.Find(&producers)
	context.JSON(http.StatusOK, producers)
}
