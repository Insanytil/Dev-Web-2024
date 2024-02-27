package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"my-app/db"
	"my-app/web"
	"os"
)

// @title local eat API
// @version 1.0
// @description This is a sample server local eat API server.
// @host localhost:8080
// @BasePath /
// @schemes http
// @produce json
func main() {
	mysql, err := sql.Open("mysql", dataSource())
	if err != nil {
		log.Fatal(err)
	}
	defer mysql.Close()
	// CORS is enabled only in prod profile
	cors := os.Getenv("profile") == "prod"
	app := web.NewApp(db.NewDB(mysql), cors)
	log.Println("Error", app)
}

func dataSource() string {
	host := "localhost"
	pass := "pass"
	if os.Getenv("profile") == "prod" {
		host = "db"
		pass = os.Getenv("db_pass")
	}
	return "goxygen:" + pass + "@tcp(" + host + ":3306)/goxygen"
}
