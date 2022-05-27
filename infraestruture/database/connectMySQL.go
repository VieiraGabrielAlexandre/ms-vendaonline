package database

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "root:vendaonline@tcp(172.18.0.3:3306)/ms-vendaonline?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("DB_USER") +
		":" +
		os.Getenv("DB_PASSWORD") +
		"@tcp(" +
		os.Getenv("DB_HOST") +
		":" +
		os.Getenv("DB_PORT") +
		")/" +
		os.Getenv("DB_DATABASE") +
		"?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	// Create from map

	db.AutoMigrate(&model.Product{}, &model.Subcategory{}, &model.SubcategoriesProducts{}, &model.Category{}, &model.Image{})

	DB = db
}
