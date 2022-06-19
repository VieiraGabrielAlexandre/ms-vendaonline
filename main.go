package main

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/configs"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/routes"
)

// @title API - Xplosão - Products and correlated
// @version 1.0
// @description This is doc for API products and co-related.

// @contact.name API Support
// @contact.url https://www.xplosao.com.br/support
// @contact.email suporte@xplosao.com.br

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	configs.Environment()
	database.Connect()
	routes.HandleRequests()
}
