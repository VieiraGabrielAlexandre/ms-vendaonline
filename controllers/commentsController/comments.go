package commentsController

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Create(c *gin.Context) {
	var comment model.Comments

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		os.Exit(1)
	}

	if err := model.ValidatorComment(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		os.Exit(1)
	}

	database.DB.Create(&comment)

	c.JSON(200, comment)
}

func Index(c *gin.Context) {
	var comments []model.Comments

	id := c.Params.ByName("id_product")
	database.DB.Where("id_product = ?", id).Find(&comments)

	c.JSON(200, comments)
}
