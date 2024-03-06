package main

import (
	"database/sql"
	"log"
	"local_eat/api/db"
	"local_eat/api/web"
	"os"

	_ "github.com/go-sql-driver/mysql"
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
	appErr := web.NewApp(db.NewDB(mysql), cors)
	log.Println("Error", appErr)
}

func dataSource() string {
	host := "localhost"
	pass := "pass"
	if os.Getenv("profile") == "prod" {
		host = "db"
		pass = os.Getenv("db_pass")
	}
	return "goxygen:" + pass + "@tcp(" + host + ":3906)/goxygen"
}
