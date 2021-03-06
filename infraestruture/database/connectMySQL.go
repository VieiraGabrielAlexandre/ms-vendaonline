package database

import (
	"fmt"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var (
	DB *gorm.DB
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panic(err)
	}

	// Create from map

	fmt.Println("Executing migrations ...")
	db.AutoMigrate(&model.Product{}, &model.Subcategory{}, &model.SubcategoriesProducts{}, &model.Category{}, &model.Image{}, &model.Comments{}, &model.Prices{}, &model.Users{})
	fmt.Println("Sucess")

	DB = db
}
