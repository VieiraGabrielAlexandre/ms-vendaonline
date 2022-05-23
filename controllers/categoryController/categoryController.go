package categoryController

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/core/domain/services/categoryService"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Show(c *gin.Context) {
	var categories []model.Category

	database.DB.Find(&categories)
	c.JSON(200, categories)
}

func Create(c *gin.Context) {
	var category model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		os.Exit(1)
	}

	if err := model.ValidatorCategory(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		os.Exit(1)
	}

	result := categoryService.Create(category)

	c.JSON(200, result)
}
