package main

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/routes"
)

func main() {
	database.Connect()
	routes.HandleRequests()
}
