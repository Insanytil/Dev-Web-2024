package web

import (
	"log"
	"my-app/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	_ "my-app/docs"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type App struct {
	d db.DB
}

func NewApp(d db.DB, corsBool bool) error {
	router := gin.Default()
	app := App{d: d}
	techHandler := app.GetTechnologies
	pingHandler := app.GetPing
	if !corsBool {
		router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
	}))
	}
	router.GET("/api/technologies", techHandler)
	router.GET("/api/ping", pingHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // The url pointing to API definition
	log.Println("Web server is available on port 8080")
	return router.Run(":8080")
}

// swagger:operation GET /api/technologies Technologies GetTechnologiesRequest
// GET Technologies
// @Summary Get all technologies
// @Description Get all technologies
// @Tags Technologies
// @Produce json
// @Success 200 {array} string
// @Router /api/technologies [get]
func (app *App) GetTechnologies(context *gin.Context) {
	technologies, err := app.d.GetTechnologies()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, technologies)
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
	technologies, err := app.d.GetTechnologies()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, technologies)
}