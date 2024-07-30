package server

import (
	"context"
	"fmt"

	database "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driven"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driven/repositories/item"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/config"
	itemHandler "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/http/item"
	services "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/services/item"
	"github.com/labstack/echo/v4"
)

func newApp(cfg config.Config) *echo.Echo {
	app := echo.New()

	database.ConectaDB(cfg.DatabaseConfig.Host, cfg.DatabaseConfig.User, cfg.DatabaseConfig.Password, cfg.DatabaseConfig.DbName, cfg.DatabaseConfig.Port)

	itemRepository := item.NewRepository(database.DB)
	itemService := services.NewItemService(itemRepository)
	itemHandler := itemHandler.NewItemHandler(itemService)
	itemHandler.RegisterRoutes(app)

	return app
}

func Start(cfg config.Config) {
	fmt.Println(context.Background(), fmt.Sprintf("Starting a server at http://%s", cfg.ServerHost))
	app := newApp(cfg)
	app.Logger.Fatal(app.Start(cfg.ServerHost))
}
