package categoryService

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
)

func Create(category model.Category) model.Category {

	database.DB.FirstOrCreate(&category, "title = ?", category.Title)
	return category
}
