package main

import (
	"fmt"
	"github.com/8soat-grupo35/tech-challenge-fase1/src/api/app"
	"github.com/8soat-grupo35/tech-challenge-fase1/src/external"
)

func main() {
	fmt.Println("Iniciado o servidor Rest com GO")
	cfg := external.GetConfig()
	app.Start(cfg)
}
