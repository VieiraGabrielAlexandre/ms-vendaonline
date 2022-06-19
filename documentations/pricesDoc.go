package documentations

import "github.com/gin-gonic/gin"

// @BasePath /api/v1
// @Security ApiKeyAuth


// Index godoc
// @Summary prices example
// @Description Prices per product
// @Tags Prices
// @Accept json
// @Param id_product path int false "Prices por id product"
// @Param Authorization header string false "Header for request"
// @Produce json
// @Success 200 {string} Index
// @Router /prices/{id_product} [get]

func Index(c *gin.Context) {

}
