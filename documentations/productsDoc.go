package documentations

import "github.com/gin-gonic/gin"

// @BasePath /api/v1

// Show godoc
// @ID show-products
// @Summary products
// @Schemes
// @Description Products per page and page limit
// @Param page query int false "Page of response products"
// @Param page_size query int false "Items per page"
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} array
// @Router /products [get]
func Show(c *gin.Context) {

}

// Index godoc
// @ID index-products
// @Summary products
// @Schemes
// @Description Product per id
// @Param id_product path int false "ID product"
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} array
// @Router /products/{id_product} [get]
func Index(c *gin.Context) {
	
}
