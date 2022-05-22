package productsService

import (
	"fmt"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
)

type Product string

func Create(product model.Product) {
	database.DB.Create(&product)

	fmt.Println(&product)
}
