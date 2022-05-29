package model

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	ID_Product int    `json:"id_product" validate:"nonzero"`
	ID_Client  int    `json:"id_client" validate:"nonzero"`
	Title      string `json:"title" validate:"nonzero"`
	Comment    string `json:"comment" validate:"nonzero"`
}

func ValidatorComment(comment *Comments) error {
	if err := validator.Validate(comment); err != nil {
		return err
	}

	return nil
}
