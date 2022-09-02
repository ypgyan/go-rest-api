package main

import (
	"fmt"
	"github.com/ypgyan/go-rest-api/database"
	"github.com/ypgyan/go-rest-api/models"
	"github.com/ypgyan/go-rest-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{1, "Nome 1", "Uma história bem legal 1"},
		{2, "Nome 2", "Uma história bem legal 2"},
	}
	database.ConnectDB()
	fmt.Println("Iniciando servidor Rest com GoLang")
	routes.HandleRequest()
}
