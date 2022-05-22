package productsController

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/core/domain/services/productsService"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Show(c *gin.Context) {
	var products []model.Product

	database.DB.Find(&products)

	c.JSON(200, products)
}

func Create(c *gin.Context) {
	var product model.Product
	productsService.Create(product)

	c.JSON(http.StatusOK, product)
}
