package productSubcategoryService

import (
	"fmt"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
)

type SubcategoryProduct string

func Create(subcategoryProduct model.SubcategoryProducts) {
	database.DB.Create(&subcategoryProduct)

	fmt.Println(&subcategoryProduct)
}
