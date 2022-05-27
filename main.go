package main

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/configs"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/routes"
)

func main() {
	configs.Environment()
	database.Connect()
	routes.HandleRequests()
}
