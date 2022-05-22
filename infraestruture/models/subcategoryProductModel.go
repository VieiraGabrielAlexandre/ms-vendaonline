package model

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type SubcategoryProduct struct {
	gorm.Model
	ID_Product     int `json:"id_product" validate:"nonzero"`
	ID_Subcategory int `json:"id_subcategory" validate:"nonzero"`
}

func ValidatorSubcategoryProduct(subcategoryProduct *SubcategoryProduct) error {
	if err := validator.Validate(subcategoryProduct); err != nil {
		return err
	}

	return nil
}
