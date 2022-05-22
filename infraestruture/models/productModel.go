package model

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string `json:"title" validate:"nonzero"`
	Brand       string `json:"brand" validate:"nonzero"`
	Collor      string `json:"collor" validate:"nonzero"`
	Description string `json:"description" validate:"nonzero"`
	Image       string `json:"image" validate:"nonzero"`
	Category    string `json:"category" validate:"nonzero"`
	Seller_id   int    `json:"seller_id" validate:"nonzero"`
}

var Products []Product

func ValidatorProduct(product *Product) error {
	if err := validator.Validate(product); err != nil {
		return err
	}

	return nil
}
