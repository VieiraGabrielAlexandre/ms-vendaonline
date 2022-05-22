package routes

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/productsController"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/subcategoryController"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/subcategoryProductsController"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/products", productsController.Show)
		v1.POST("/products", productsController.Create)

		v1.GET("/subcategory", subcategoryController.Show)
		v1.POST("/subcategory", subcategoryController.Create)

		v1.GET("/subcategory-products", subcategoryProductsController.Show)
		v1.POST("/subcategory-products", subcategoryProductsController.Create)
	}

	router.Run(":5566")
}
