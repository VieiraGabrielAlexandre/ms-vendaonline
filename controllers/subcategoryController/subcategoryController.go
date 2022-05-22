package subcategoryController

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Show(c *gin.Context) {
	var subcategories []model.Subcategory

	database.DB.Find(&subcategories)

	c.JSON(200, subcategories)
}

func Create(c *gin.Context) {
	var subcategory model.Subcategory

	if err := c.ShouldBindJSON(&subcategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := model.ValidatorSubcategory(&subcategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&subcategory)
	c.JSON(http.StatusOK, subcategory)
}
