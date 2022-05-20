package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Titulo       string `json:"titulo"`
	Marca        string `json:"marca"`
	Cor          string `json:"cor"`
	Descricao    string `json:"descricao"`
	Imagem       string `json:"imagem"`
	Categoria    string `json:"categoria"`
	Vendendor_id int    `json:"vendendor_id"`
}
