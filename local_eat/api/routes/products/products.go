package products

import (
	"local_eat/api/initializers"
	"local_eat/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	users := route.Group("/api/products")
	{
		users.GET("", GetProducts)
	}
}

func GetProducts(context *gin.Context) {
	var products []models.Product
	initializers.DB.Find(&products)
	context.JSON(http.StatusOK, products)
}
