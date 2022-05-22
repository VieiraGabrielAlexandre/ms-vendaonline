package subcategoryProductsController

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	var subcategoryProducts []model.SubcategoryProduct

	database.DB.Find(&subcategoryProducts)

	c.JSON(200, subcategoryProducts)
}

func Create(c *gin.Context) {
	var subcategoryProducts []model.SubcategoryProduct
	
	c.JSON(200, subcategoryProducts)

}
