package server

import (
	"context"
	"fmt"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/config"
	"github.com/labstack/echo/v4"
)

func newApp(cfg config.Config) *echo.Echo {
	app := echo.New()

	//database.ConectaDB()
	//
	//itemRepository := item.NewRepository(database.DB)
	//itemService := itemService.New(itemRepository)

	return app
}

func Start(cfg config.Config) {
	fmt.Println(context.Background(), fmt.Sprintf("Starting a server at http://%s", cfg.ServerHost))

	app := newApp(cfg)

	app.Logger.Fatal(app.Start(cfg.ServerHost))
}
