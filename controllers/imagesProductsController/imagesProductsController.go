package imagesProductsController

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var images []model.Image

	id := c.Params.ByName("id_product")
	database.DB.Find(&images, "id_product = ?", id)

	c.JSON(200, images)
}
