package initializers

import (
	"log"
	"os"

	model "local_eat/api/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectBD() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	addr := os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	param := "?parseTime=True"
	dsn := user + ":" + pass + "@tcp(" + addr + ")/" + dbName + param

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(DB)
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func SyncDB() {
	err := DB.AutoMigrate(&model.Users{}, &model.Producers{}, &model.Category{}, &model.Product{},
		&model.Company{}, &model.RelCompProd{})
	if err != nil {
		log.Fatal(err)
	}
	// Définir les contraintes de clé étrangère supplémentaires si nécessaire
	// Par exemple, pour MySQL, vous pouvez utiliser DB.Exec() pour exécuter des requêtes SQL
	// pour définir les contraintes de clé étrangère.
	// Voici un exemple de requête SQL pour définir une contrainte de clé étrangère dans MySQL :
	// _, err = DB.Exec("ALTER TABLE producers ADD CONSTRAINT fk_users FOREIGN KEY (username) REFERENCES users(username) ON UPDATE CASCADE ON DELETE SET NULL;")
	// if err != nil {
	//     log.Fatal(err)
	// }

	// Répétez ce processus pour chaque clé étrangère que vous souhaitez définir.
}
