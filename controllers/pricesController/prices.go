package pricesController

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var prices model.Prices

	id := c.Params.ByName("id_product")

	database.DB.First(&prices).Where("id_product = ?", id)
	c.JSON(200, prices)
}
