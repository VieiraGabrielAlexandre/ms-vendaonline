package routes

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/categoryController"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/imagesProductsController"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/productsController"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/relatedProductsController"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/subcategoryController"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/subcategoryProductsController"
	"github.com/gin-gonic/gin"
	"os"
)

func HandleRequests() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		//Products
		v1.GET("/products", productsController.Show)
		v1.POST("/products", productsController.Create)
		v1.GET("/products/:id", productsController.Index)

		//RelatedProducts
		v1.GET("/related-products/:category", relatedProductsController.Show)

		//Subcategory
		v1.GET("/subcategory", subcategoryController.Show)
		v1.POST("/subcategory", subcategoryController.Create)

		//Subcategory Products
		v1.GET("/subcategory-products", subcategoryProductsController.Show)
		v1.POST("/subcategory-products", subcategoryProductsController.Create)

		//Category
		v1.GET("/category", categoryController.Show)
		v1.POST("/category", categoryController.Create)

		//Images
		v1.GET("/images/:id_product", imagesProductsController.Index)
	}

	router.Run(":" + os.Getenv("SERVER_PORT"))

}
