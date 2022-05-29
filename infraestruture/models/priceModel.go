package model

import "gorm.io/gorm"

type Prices struct {
	gorm.Model
	ID_Product int     `json:"int" validate:"nonzero"`
	Price      float64 `json:"price" validate:"nonzero"`
	Full_Price float64 `json:"full_price" validate:"nonzero"`
}
