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
	docs "github.com/VieiraGabrielAlexandre/ms-vendaonline/docs"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/middlewares"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

func HandleRequests() {
	router := gin.Default()

	authMiddleware := middlewares.Auth()

	docs.SwaggerInfo.BasePath = "/api/v1"

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router.POST("/login", authMiddleware.LoginHandler)

	auth := router.Group("/api/v1")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		//Products
		auth.GET("/products", productsController.Show)
		auth.POST("/products", productsController.Create)
		auth.GET("/products/:id", productsController.Index)

		//Prices
		auth.GET("/prices/:id_product", pricesController.Index)

		//RelatedProducts
		auth.GET("/related-products/:category", relatedProductsController.Show)

		//Subcategory
		auth.GET("/subcategory", subcategoryController.Show)
		auth.POST("/subcategory", subcategoryController.Create)

		//Subcategory Products
		auth.GET("/subcategory-products", subcategoryProductsController.Show)
		auth.POST("/subcategory-products", subcategoryProductsController.Create)

		//Category
		auth.POST("/category", categoryController.Create)
		auth.GET("/category", categoryController.Show)

		//Images
		auth.GET("/images/:id_product", imagesProductsController.Index)

		//Coments Produts
		auth.POST("comments", commentsController.Create)
		auth.GET("comments/:id_product", commentsController.Index)

		//Stars Products
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":" + os.Getenv("SERVER_PORT"))

}
