package products

import (
	"net/http"

	"local_eat/api/initializers"
	"local_eat/api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(route *gin.Engine) {
	products := route.Group("/api/products")
	{
		products.GET("", GetProducts)
		products.POST("add-product", CreateProduct)
		products.GET("categories", GetCategories)
		products.GET("by-category", GetProductsByCategory) // New route
	}
}

func GetCategories(context *gin.Context) {
	var categories []models.Category
	result := initializers.DB.Find(&categories)
	if result.Error == gorm.ErrRecordNotFound {
		context.JSON(http.StatusNotFound, gin.H{})
		return
	}
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusOK, categories)

}

// @Summary Get products by category
// @Description Get products filtered by category id
// @Tags Products
// @Produce json
// @Param categoryId query string true "Category ID"
// @Success 200 {array} models.Product
// @Failure 404 "Not found"
// @Failure 500 "Internal server error"
// @Router /products/by-category [get]
func GetProductsByCategory(context *gin.Context) {
	categoryId := context.Query("categoryId")
	if categoryId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Category ID is required"})
		return
	}

	var products []models.Product
	result := initializers.DB.Where("category_id = ?", categoryId).Find(&products)
	if result.Error == gorm.ErrRecordNotFound {
		context.JSON(http.StatusNotFound, gin.H{})
		return
	}
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusOK, products)
}

// @Summary Get products
// @Description Get products id, name, picture, category and description
// @Tags Products
// @Produce json
// @Success 200 {array} models.Product
// @Failure 404 "Not found"
// @Failure 500 "Internal server error"
// @Router /products [get]
func GetProducts(context *gin.Context) {
	var products []models.Product
	result := initializers.DB.Find(&products)
	if result.Error == gorm.ErrRecordNotFound {
		context.JSON(http.StatusNotFound, gin.H{})
		return

	}
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusOK, products)
}

// @Summary Create product
// @Description Create a new product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product"
// @Success 200 "Product created"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /products [post]
func CreateProduct(context *gin.Context) {
	var product models.Product
	err := context.ShouldBindJSON(&product)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	result := initializers.DB.Create(&product)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Product created"})
}
