package main

import (
	"fmt"

	database "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driven"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/http"
)

func main() {
	fmt.Println("Iniciado o servidor Rest com GO")

	database.ConectaDB()

	http.SetupRouter(database.DB)
}
