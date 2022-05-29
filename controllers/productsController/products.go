package productsController

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/core/domain/services/productsService"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/scopes"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Show(c *gin.Context) {
	var products []model.Product

	database.DB.Scopes(scopes.Paginate(c)).
		Preload("Images").
		Preload("Subcategories").
		Preload("Prices").
		Find(&products)

	c.JSON(200, products)
}

func Create(c *gin.Context) {
	var product model.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		os.Exit(1)
	}

	if err := model.ValidatorProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		os.Exit(1)
	}

	result := productsService.Create(product)

	c.JSON(http.StatusOK, result)
}

func Index(c *gin.Context) {
	var product model.Product

	id := c.Params.ByName("id")

	database.DB.Model(&product).
		Preload("Images").
		Preload("Subcategories").
		Preload("Prices").
		First(&product, "id = ?", id)

	c.JSON(200, product)
}
