package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name     string `json:"name" validate:"nonzero"`
	Email    string `json:"email" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
	Active   bool   `json:"active" validate:"nonzero"`
}
