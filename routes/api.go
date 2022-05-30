package routes

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/categoryController"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/commentsController"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/imagesProductsController"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/pricesController"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/productsController"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/relatedProductsController"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/subcategoryController"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers/subcategoryProductsController"
	swaggerfiles "github.com/swaggo/files"

	docs "github.com/VieiraGabrielAlexandre/ms-vendaonline/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

func HandleRequests() {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group("/api/v1")
	{
		//Products
		v1.GET("/products", productsController.Show)
		v1.POST("/products", productsController.Create)
		v1.GET("/products/:id", productsController.Index)

		//Prices
		v1.GET("/prices/:id_product", pricesController.Index)

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

		//Coments Produts
		v1.POST("comments", commentsController.Create)
		v1.GET("comments/:id_product", commentsController.Index)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":" + os.Getenv("SERVER_PORT"))

}
