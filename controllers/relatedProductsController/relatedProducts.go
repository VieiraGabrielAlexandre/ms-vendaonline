package relatedProductsController

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	var product []model.Product

	category := c.Params.ByName("category")

	database.DB.Model(&product).
		Preload("Images").
		Preload("Subcategories").
		Find(&product, "category = ?", category).
		Limit(5)

	c.JSON(200, product)
}
