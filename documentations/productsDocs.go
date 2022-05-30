package documentations

import "github.com/gin-gonic/gin"

// @BasePath /api/v1

// Show godoc
// @ID show-products
// @Summary products
// @Schemes
// @Description Products por page and page limit
// @Param page query int false "Page of response products"
// @Param page_size query int false "Items per page"
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} array
// @Router /products [get]
func Show(c *gin.Context) {

}
