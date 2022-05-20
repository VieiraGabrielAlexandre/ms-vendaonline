package routes

import (
	produtos "github.com/VieiraGabrielAlexandre/ms-vendaonline/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/products", produtos.Show)
	r.POST("/products", produtos.Create)

	r.Run(":5566")
}
