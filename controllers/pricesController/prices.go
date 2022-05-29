package pricesController

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var prices model.Prices

	database.DB.First(&prices)
	c.JSON(200, prices)
}
