package subcategoryService

import (
	"fmt"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"os"
)

type Product string

func Create(subcategoryProduct []model.SubcategoryProduct) {
	fmt.Println(subcategoryProduct)
	os.Exit(1)
	database.DB.Create(&subcategoryProduct)

	fmt.Println(&subcategoryProduct)
}
