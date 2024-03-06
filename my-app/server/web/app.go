package web

import (
	"log"
	"net/http"
	"local_eat/api/db"
	_ "local_eat/api/docs"
	
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type App struct {
	db db.DB
}

func NewApp(d db.DB, corsBool bool) error {
	router := gin.Default()
	app := App{db: d}
	if !corsBool {
		router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
	}))
	}
	router.GET("/api/ping", app.GetPing)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // The url pointing to API definition
	log.Println("Web server is available on port 8080")
	return router.Run(":8080")
}

// swagger:operation GET /api/ping Ping GetPingRequest
// GET Ping
// @Summary Check API status
// @Description Check API status
// @Tags Ping
// @Produce json
// @Success 200 {array} string
// @Router /api/ping [get]
func (app *App) GetPing(context *gin.Context) {
	technologies, err := app.db.GetTechnologies()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, technologies)
}