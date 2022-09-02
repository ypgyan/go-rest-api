package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	connectStr := "host=localhost user=root password=root dbname=root port=54032 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(connectStr))
	if err != nil {
		log.Panicln("Erro ao conectar com o banco de dados")
	}
}
