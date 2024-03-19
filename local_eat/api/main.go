package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "local_eat/api/docs"
	"local_eat/api/initializers"
	"local_eat/api/routes/auth"
	"local_eat/api/routes/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectBD()
	initializers.SyncDB()
}

// @title local eat API
// @version 1.0
// @description This is a sample server local eat API server.
// @host localhost:8080
// @BasePath /
// @schemes http
// @produce json
func main() {
	// CORS is enabled only in prod profile
	router := gin.Default()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if !(os.Getenv("profile") == "prod") {
		router.Use(cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
			AllowHeaders: []string{"*"},
		}))
	}
	users.Routes(router)
	auth.Routes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // The url pointing to API definition
	log.Println("Web server is available on port 8080")
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1) // Add buffer size of 1
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
