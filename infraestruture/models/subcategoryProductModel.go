package model

import (
	"gorm.io/gorm"
)

type SubcategoriesProducts struct {
	gorm.Model
	ID_Product     int `json:"id_product" validate:"nonzero"`
	ID_Subcategory int `json:"id_subcategory" validate:"nonzero"`
}

type SubcategoryProducts []SubcategoriesProducts
