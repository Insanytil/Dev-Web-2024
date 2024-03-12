package web

import (
	"log"
	"local_eat/api/db"
	"local_eat/api/routes/users"
	_ "local_eat/api/docs"
	
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func NewApp(db db.DB, corsBool bool) error {
	router := gin.Default()
	if !corsBool {
		router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
	}))
	}
	users.Routes(router, db)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // The url pointing to API definition
	log.Println("Web server is available on port 8080")
	return router.Run(":8080")
}