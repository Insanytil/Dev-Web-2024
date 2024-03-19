package main

import (
	"database/sql"
	"local_eat/api/web"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// @title local eat API
// @version 1.0
// @description This is a sample server local eat API server.
// @host localhost:8080
// @BasePath /
// @schemes http
// @produce json
func main() {
	mysql, mysqlErr := sql.Open("mysql", dataSource())
	if mysqlErr != nil {
		log.Fatal(mysqlErr)
	}
	defer mysql.Close()
	// CORS is enabled only in prod profile
	cors := os.Getenv("profile") == "prod"
	web.NewApp(mysql, cors)
}

func dataSource() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := mysql.Config{
        User:   os.Getenv("DB_USER"),
        Passwd: os.Getenv("DB_PASS"),
        Net:    "tcp",
        Addr:   os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
        DBName: os.Getenv("DB_NAME"),
    }
	return cfg.FormatDSN()
}
