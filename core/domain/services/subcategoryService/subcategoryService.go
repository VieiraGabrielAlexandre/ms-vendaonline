package subcategoryService

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
)

type Product string

func Create(subcategory model.Subcategory) model.Subcategory {
	database.DB.FirstOrCreate(&subcategory, "title = ?", subcategory.Title)
	return subcategory
}
