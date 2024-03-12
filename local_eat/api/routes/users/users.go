package users

import (
	"local_eat/api/db"
	"local_eat/api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine, d db.DB) {
	users := route.Group("/api/producers")
	{
		users.Use(middleware.DBMiddleware(d))
		users.GET("/", GetProducers)
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
	db := context.MustGet("db").(db.DB)
	producers, err := db.GetProducers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	context.JSON(http.StatusOK, producers)
}
