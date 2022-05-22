package model

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Subcategory struct {
	gorm.Model
	Title string `json:"title" validate:"nonzero"`
}

func ValidatorSubcategory(subcategory *Subcategory) error {
	if err := validator.Validate(subcategory); err != nil {
		return err
	}

	return nil
}
