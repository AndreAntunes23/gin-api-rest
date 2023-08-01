package main

import (
	"github.com/AndreAntunes23/gin-api-rest/database"
	"github.com/AndreAntunes23/gin-api-rest/routes"
)

func main() {
	database.ConectaDB()
	routes.HandleRequests()
}
