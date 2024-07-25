package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaDB() {
	conexao := "host=postgres user=root password=root dbname=root port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(conexao))

	if err != nil {
		log.Panic("Erro na conexao com banco de dados")
	}

	DB = db
}
