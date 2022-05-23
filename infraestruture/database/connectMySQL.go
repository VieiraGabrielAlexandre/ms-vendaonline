package database

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:vendaonline@tcp(127.0.0.1:3307)/ms-vendaonline?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Erro ao conectar ao banco")
	}

	// Create from map

	db.AutoMigrate(&model.Product{}, &model.Subcategory{}, &model.SubcategoriesProducts{}, &model.Category{}, &model.Image{})

	DB = db
}
