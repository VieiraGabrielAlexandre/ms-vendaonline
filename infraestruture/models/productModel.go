package model

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Title         string                 `json:"title" validate:"nonzero"`
	Brand         string                 `json:"brand" validate:"nonzero"`
	Collor        string                 `json:"collor" validate:"nonzero"`
	Description   string                 `json:"description" validate:"nonzero"`
	Image         string                 `json:"image" validate:"nonzero"`
	Category      string                 `json:"category" validate:"nonzero"`
	Seller_id     int                    `json:"seller_id" validate:"nonzero"`
	Images        []Images               `gorm:"foreignKey:ID_Product;constraint:onDelete:SET NULL,onUpdate:CASCADE" validate:"nonzero"`
	Subcategories []SubcategoriesProduct `gorm:"foreignKey:ID_Product;constraint:onDelete:SET NULL,onUpdate:CASCADE" validate:"nonzero"`
	Prices        Price                  `gorm:"foreignKey:ID_Product;constraint:onDelete:SET NULL,onUpdate:CASCADE" validate:"nonzero"`
}

type Images struct {
	gorm.Model
	Title      string `json:"title" validate:"nonzero"`
	URL        string `json:"url" validate:"nonzero"`
	ID_Product int    `json:"id_product"`
}

type SubcategoriesProduct struct {
	gorm.Model
	ID_Product     int `json:"id_product"`
	ID_Subcategory int `json:"id_subcategory" validate:"nonzero"`
}

type Price struct {
	gorm.Model
	ID_Product int     `json:"int"`
	Price      float64 `json:"price" validate:"nonzero"`
	Full_Price float64 `json:"full_price" validate:"nonzero"`
}

func ValidatorProduct(product *Product) error {
	if err := validator.Validate(product); err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
