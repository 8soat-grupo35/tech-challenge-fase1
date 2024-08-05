package server

import (
	"context"
	"fmt"

	database "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driven"
	clientRepository "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driven/repositories/client"
	itemRepository "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driven/repositories/item"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/config"
	clientHandler "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/http/client"
	itemHandler "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/http/item"
	clientService "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/services/client"
	itemService "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/services/item"
	"github.com/labstack/echo/v4"
)

func newApp(cfg config.Config) *echo.Echo {
	app := echo.New()

	database.ConectaDB(cfg.DatabaseConfig.Host, cfg.DatabaseConfig.User, cfg.DatabaseConfig.Password, cfg.DatabaseConfig.DbName, cfg.DatabaseConfig.Port)

	itemRepository := itemRepository.NewRepository(database.DB)
	itemService := itemService.NewItemService(itemRepository)
	itemHandler := itemHandler.NewItemHandler(itemService)
	itemHandler.RegisterRoutes(app)

	clientRepository := clientRepository.NewRepository(database.DB)
	clientService := clientService.NewClientService(clientRepository)
	clientHandler := clientHandler.NewClientHandler(clientService)
	clientHandler.RegisterRoutes(app)

	return app
}

func Start(cfg config.Config) {
	fmt.Println(context.Background(), fmt.Sprintf("Starting a server at http://%s", cfg.ServerHost))
	app := newApp(cfg)
	app.Logger.Fatal(app.Start(cfg.ServerHost))
}
