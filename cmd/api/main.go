package main

import (
	"fmt"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/config"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/http/server"
)

func main() {
	fmt.Println("Iniciado o servidor Rest com GO")
	cfg := config.GetConfig()
	server.Start(cfg)
}
