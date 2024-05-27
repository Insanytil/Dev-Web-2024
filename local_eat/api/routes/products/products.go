package products

import (
	"net/http"
	"time"

	"local_eat/api/initializers"
	"local_eat/api/middleware"
	"local_eat/api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(route *gin.Engine) {
	products := route.Group("/api/products")
	{
		products.GET("", GetProducts)
		products.POST("add-product", middleware.AuthMiddleware, CreateProduct)
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
	// Retrieve the user from the context
	user, _ := context.Get("user")
	foundUser := user.(models.Users).Username

	// Find the producer based on the username
	var producer models.Producers
	result := initializers.DB.Where("username = ?", foundUser).First(&producer)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving producer from database"})
		return
	}
	if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Producer not found"})
		return
	}

	// Find the related company-producer relationship
	var relCompProd models.RelCompProd
	result2 := initializers.DB.Where("producer_id = ?", producer.ID).First(&relCompProd)
	if result2.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving CompanyName from database"})
		return
	}

	// Bind the product data from the request
	var product models.Product
	err := context.ShouldBindJSON(&product)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Bind the catalog details data from the request
	var body models.CatalogDetails
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the new product and catalog details
	newCatalogDetails := models.CatalogDetails{
		ID:           body.ID,
		CompanyName:  relCompProd.CompanyName,
		ProductId:    product.ID,
		CreatedAt:    time.Now(),
		Quantity:     body.Quantity,
		Availability: body.Quantity > 0,
		Price:        body.Price,
	}
	productResult := initializers.DB.Create(&newCatalogDetails)
	if productResult.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{})
}
