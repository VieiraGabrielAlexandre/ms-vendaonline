package productsService

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
)

type Product string

func Create(product model.Product) model.Product {
	database.DB.Create(&product)

	return product
}
