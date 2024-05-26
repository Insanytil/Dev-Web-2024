package initializers

import (
	"log"
	"os"

	"local_eat/api/models"

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
		log.Println("No .env file found")
	}
}

func SyncDB() {
	err := DB.AutoMigrate(&models.Images{}, &models.Users{}, &models.Producers{}, &models.Category{}, &models.Product{},
		&models.Company{}, &models.CatalogDetails{}, &models.RelCompProd{})
	if err != nil {
		log.Fatal(err)
	}
	initializeCategories(DB)
	disableForeignKeyChecks()
	initializeProducts(DB)
	defer enableForeignKeyChecks()
}

func initializeCategories(db *gorm.DB) {
	categories := []models.Category{
		{ID: "CAT1", Name: "Fruits", Description: stringPtr("Catégorie des fruits")},
		{ID: "CAT2", Name: "Légumes", Description: stringPtr("Catégorie des légumes")},
		{ID: "CAT3", Name: "Viandes", Description: stringPtr("Catégorie des viandes")},
		{ID: "CAT4", Name: "Poissons et Fruits de Mer", Description: stringPtr("Catégorie des poissons et fruits de mer")},
		{ID: "CAT5", Name: "Produits Laitiers", Description: stringPtr("Catégorie des produits laitiers")},
		{ID: "CAT6", Name: "Céréales et Grains", Description: stringPtr("Catégorie des céréales et grains")},
		{ID: "CAT7", Name: "Légumineuses", Description: stringPtr("Catégorie des légumineuses")},
		{ID: "CAT8", Name: "Noix et Graines", Description: stringPtr("Catégorie des noix et graines")},
		{ID: "CAT9", Name: "Herbes et Épices", Description: stringPtr("Catégorie des herbes et épices")},
		{ID: "CAT10", Name: "Pains et Pâtisseries", Description: stringPtr("Catégorie des pains et pâtisseries")},
		{ID: "CAT11", Name: "Boissons", Description: stringPtr("Catégorie des boissons")},
		{ID: "CAT12", Name: "Plats Préparés", Description: stringPtr("Catégorie des plats préparés")},
		{ID: "CAT13", Name: "Sauces et Condiments", Description: stringPtr("Catégorie des sauces et condiments")},
		{ID: "CAT14", Name: "Confiseries et Snacks", Description: stringPtr("Catégorie des confiseries et snacks")},
		{ID: "CAT15", Name: "Produits de Boulangerie", Description: stringPtr("Catégorie des produits de boulangerie")},
		{ID: "CAT16", Name: "Produits Surgelés", Description: stringPtr("Catégorie des produits surgelés")},
	}

	for _, category := range categories {
		if err := db.FirstOrCreate(&models.Category{}, models.Category{ID: category.ID}).Error; err != nil {
			log.Printf("could not create category %s: %v", category.Name, err)
		} else {
			db.Save(&category) // Save the full category with all fields
		}
	}
}
func initializeProducts(DB *gorm.DB) {
	products := []models.Product{
		{ID: "PROD1", Name: "Pomme", CategoryID: "CAT1", Description: stringPtr("Une pomme rouge fraîche."), Picture: "../images/products/apple.jpg"},
		{ID: "PROD2", Name: "Banane", CategoryID: "CAT1", Description: stringPtr("Une banane mûre."), Picture: "../images/products/banana.jpg"},
		{ID: "PROD3", Name: "Orange", CategoryID: "CAT1", Description: stringPtr("Une orange juteuse."), Picture: "../images/products/orange.jpg"},
		{ID: "PROD4", Name: "Fraise", CategoryID: "CAT1", Description: stringPtr("Des fraises sucrées."), Picture: "../images/products/strawberry.jpg"},
		{ID: "PROD5", Name: "Raisins", CategoryID: "CAT1", Description: stringPtr("Une grappe de raisins."), Picture: "../images/products/grapes.jpg"},
		{ID: "PROD6", Name: "Pastèque", CategoryID: "CAT1", Description: stringPtr("Une pastèque juteuse."), Picture: "../images/products/watermelon.jpg"},
		{ID: "PROD7", Name: "Mangue", CategoryID: "CAT1", Description: stringPtr("Une mangue tropicale."), Picture: "../images/products/mango.jpg"},
		{ID: "PROD8", Name: "Ananas", CategoryID: "CAT1", Description: stringPtr("Un ananas frais."), Picture: "../images/products/pineapple.jpg"},
		{ID: "PROD9", Name: "Carotte", CategoryID: "CAT2", Description: stringPtr("Une carotte croquante."), Picture: "../images/products/carrot.jpg"},
		{ID: "PROD10", Name: "Brocoli", CategoryID: "CAT2", Description: stringPtr("Brocoli frais."), Picture: "../images/products/broccoli.jpg"},
		{ID: "PROD11", Name: "Épinard", CategoryID: "CAT2", Description: stringPtr("Feuilles d'épinard fraîches."), Picture: "../images/products/spinach.jpg"},
		{ID: "PROD12", Name: "Tomate", CategoryID: "CAT2", Description: stringPtr("Une tomate mûre."), Picture: "../images/products/tomato.jpg"},
		{ID: "PROD13", Name: "Concombre", CategoryID: "CAT2", Description: stringPtr("Un concombre frais."), Picture: "../images/products/cucumber.jpg"},
		{ID: "PROD14", Name: "Poivron", CategoryID: "CAT2", Description: stringPtr("Un poivron coloré."), Picture: "../images/products/bell_pepper.jpg"},
		{ID: "PROD15", Name: "Laitue", CategoryID: "CAT2", Description: stringPtr("Feuilles de laitue croquantes."), Picture: "../images/products/lettuce.jpg"},
		{ID: "PROD16", Name: "Pomme de terre", CategoryID: "CAT2", Description: stringPtr("Un sac de pommes de terre."), Picture: "../images/products/potato.jpg"},
		{ID: "PROD17", Name: "Steak", CategoryID: "CAT3", Description: stringPtr("Un steak juteux."), Picture: "../images/products/steak.jpg"},
		{ID: "PROD18", Name: "Poitrine de poulet", CategoryID: "CAT3", Description: stringPtr("Poitrine de poulet désossée."), Picture: "../images/products/chicken_breast.jpg"},
		{ID: "PROD19", Name: "Côtelette de porc", CategoryID: "CAT3", Description: stringPtr("Une côtelette de porc tendre."), Picture: "../images/products/pork_chop.jpg"},
		{ID: "PROD20", Name: "Filet de saumon", CategoryID: "CAT4", Description: stringPtr("Un filet de saumon."), Picture: "../images/products/salmon_fillet.jpg"},
		{ID: "PROD21", Name: "Thon", CategoryID: "CAT4", Description: stringPtr("Steak de thon frais."), Picture: "../images/products/tuna.jpg"},
		{ID: "PROD22", Name: "Crevettes", CategoryID: "CAT4", Description: stringPtr("Crevettes fraîches."), Picture: "../images/products/shrimp.jpg"},
		{ID: "PROD23", Name: "Lait", CategoryID: "CAT5", Description: stringPtr("Un carton de lait."), Picture: "../images/products/milk.jpg"},
		{ID: "PROD24", Name: "Fromage", CategoryID: "CAT5", Description: stringPtr("Un bloc de fromage."), Picture: "../images/products/cheese.jpg"},
		{ID: "PROD25", Name: "Yaourt", CategoryID: "CAT5", Description: stringPtr("Un pot de yaourt."), Picture: "../images/products/yogurt.jpg"},
		{ID: "PROD26", Name: "Riz", CategoryID: "CAT6", Description: stringPtr("Un sac de riz."), Picture: "../images/products/rice.jpg"},
		{ID: "PROD27", Name: "Pâtes", CategoryID: "CAT6", Description: stringPtr("Un paquet de pâtes."), Picture: "../images/products/pasta.jpg"},
		{ID: "PROD28", Name: "Avoine", CategoryID: "CAT6", Description: stringPtr("Un paquet d'avoine."), Picture: "../images/products/oats.jpg"},
		{ID: "PROD29", Name: "Lentilles", CategoryID: "CAT7", Description: stringPtr("Un paquet de lentilles."), Picture: "../images/products/lentils.jpg"},
		{ID: "PROD30", Name: "Pois chiches", CategoryID: "CAT7", Description: stringPtr("Une boîte de pois chiches."), Picture: "../images/products/chickpeas.jpg"},
		{ID: "PROD31", Name: "Amandes", CategoryID: "CAT8", Description: stringPtr("Un sac d'amandes."), Picture: "../images/products/almonds.jpg"},
		{ID: "PROD32", Name: "Noix", CategoryID: "CAT8", Description: stringPtr("Un sac de noix."), Picture: "../images/products/walnuts.jpg"},
		{ID: "PROD33", Name: "Basilic", CategoryID: "CAT9", Description: stringPtr("Feuilles de basilic frais."), Picture: "../images/products/basil.jpg"},
		{ID: "PROD34", Name: "Origan", CategoryID: "CAT9", Description: stringPtr("Origan séché."), Picture: "../images/products/oregano.jpg"},
		{ID: "PROD35", Name: "Pain", CategoryID: "CAT10", Description: stringPtr("Une miche de pain."), Picture: "../images/products/bread.jpg"},
		{ID: "PROD36", Name: "Croissant", CategoryID: "CAT10", Description: stringPtr("Un croissant frais."), Picture: "../images/products/croissant.jpg"},
		{ID: "PROD37", Name: "Jus d'orange", CategoryID: "CAT11", Description: stringPtr("Une bouteille de jus d'orange."), Picture: "../images/products/orange_juice.jpg"},
		{ID: "PROD38", Name: "Jus de pomme", CategoryID: "CAT11", Description: stringPtr("Une bouteille de jus de pomme."), Picture: "../images/products/apple_juice.jpg"},
		{ID: "PROD39", Name: "Soda", CategoryID: "CAT11", Description: stringPtr("Une canette de soda."), Picture: "../images/products/soda.jpg"},
		{ID: "PROD40", Name: "Lasagne", CategoryID: "CAT12", Description: stringPtr("Un plateau de lasagne."), Picture: "../images/products/lasagna.jpg"},
		{ID: "PROD41", Name: "Pizza surgelée", CategoryID: "CAT12", Description: stringPtr("Une pizza surgelée."), Picture: "../images/products/frozen_pizza.jpg"},
		{ID: "PROD42", Name: "Ketchup", CategoryID: "CAT13", Description: stringPtr("Une bouteille de ketchup."), Picture: "../images/products/ketchup.jpg"},
		{ID: "PROD43", Name: "Moutarde", CategoryID: "CAT13", Description: stringPtr("Une bouteille de moutarde."), Picture: "../images/products/mustard.jpg"},
		{ID: "PROD44", Name: "Barre chocolatée", CategoryID: "CAT14", Description: stringPtr("Une barre chocolatée."), Picture: "../images/products/chocolate_bar.jpg"},
		{ID: "PROD45", Name: "Bonbons", CategoryID: "CAT14", Description: stringPtr("Un paquet de bonbons."), Picture: "../images/products/candy.jpg"},
		{ID: "PROD46", Name: "Gâteau", CategoryID: "CAT15", Description: stringPtr("Une part de gâteau."), Picture: "../images/products/cake.jpg"},
		{ID: "PROD47", Name: "Légumes surgelés", CategoryID: "CAT16", Description: stringPtr("Un paquet de légumes surgelés."), Picture: "../images/products/frozen_vegetables.jpg"},
		{ID: "PROD48", Name: "Crème glacée", CategoryID: "CAT16", Description: stringPtr("Un pot de crème glacée."), Picture: "../images/products/ice_cream.jpg"},
		{ID: "PROD49", Name: "Baies surgelées", CategoryID: "CAT16", Description: stringPtr("Un paquet de baies surgelées."), Picture: "../images/products/frozen_berries.jpg"},
		{ID: "PROD50", Name: "Viande surgelée", CategoryID: "CAT16", Description: stringPtr("Un paquet de viande surgelée."), Picture: "../images/products/frozen_meat.jpg"},
	}

	for _, product := range products {
		if err := DB.FirstOrCreate(&models.Product{}, models.Product{ID: product.ID}).Error; err != nil {
			log.Printf("could not create product %s: %v", product.Name, err)
		} else {
			DB.Save(&product) // Save the full product with all fields
		}
	}
}

// Helper function to create a string pointer from a string literal
func stringPtr(s string) *string {
	return &s
}

func disableForeignKeyChecks() {
	DB.Exec("SET FOREIGN_KEY_CHECKS=0;")
}
func enableForeignKeyChecks() {
	DB.Exec("SET FOREIGN_KEY_CHECKS=1;")
}
