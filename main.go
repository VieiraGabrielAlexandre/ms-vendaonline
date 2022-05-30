package main

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/configs"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/routes"
)

// @title Swagger Example API
// @version 1.0
// @description This is doc for API products and co-related.

// @contact.name API Support
// @contact.url https://www.xplosap.com.br/support
// @contact.email suporte@xplosao.com.br

// @BasePath /v1/api
func main() {
	configs.Environment()
	database.Connect()
	routes.HandleRequests()
}
