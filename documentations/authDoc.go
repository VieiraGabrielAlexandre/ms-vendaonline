package documentations

import "github.com/gin-gonic/gin"

// Index godoc
// @Summary authorization example
// @Description Request to generate a JWT token to your requests
// @Tags Authorization
// @Accept json
// @Param username query string false "User name"
// @Param password query string false "Password"
// @Produce json
// @Success 200 {string} Index
// @Router /login [post]
func Index(c *gin.Context) {

}
