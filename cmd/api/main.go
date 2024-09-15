package main

import (
	"fmt"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/api/server"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/external"
)

func main() {
	fmt.Println("Iniciado o servidor Rest com GO")
	cfg := external.GetConfig()
	server.Start(cfg)
}
