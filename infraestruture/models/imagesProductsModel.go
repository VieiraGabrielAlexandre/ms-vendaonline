package model

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Title      string `json:"title" validate:"nonzero"`
	URL        string `json:"url" validate:"nonzero"`
	ID_Product int    `json:"id_product" validate:"nonzero"`
}

func ValidatorImage(image *Category) error {
	if err := validator.Validate(image); err != nil {
		return err
	}
	return nil
}
