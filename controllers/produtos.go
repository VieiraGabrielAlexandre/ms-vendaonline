package produtos

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Show(c *gin.Context) {
	c.JSON(200, gin.H{
		"API diz": "Olá teste",
	})
}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&product)
	c.JSON(http.StatusOK, product)
}
