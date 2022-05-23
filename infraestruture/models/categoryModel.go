package model

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title string `json:"title" validate:"nonzero"`
}

func ValidatorCategory(category *Category) error {
	if err := validator.Validate(category); err != nil {
		return err
	}
	return nil
}
