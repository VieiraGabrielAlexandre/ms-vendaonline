package documentations

import "github.com/gin-gonic/gin"

// @BasePath /api/v1

// Index godoc
// @Summary prices example
// @Description Prices per product
// @Tags Prices
// @Accept json
// @Param id_product path int false "Prices por id product"
// @Produce json
// @Success 200 {string} Index
// @Router /prices/{id_product} [get]
func Index(c *gin.Context) {

}
