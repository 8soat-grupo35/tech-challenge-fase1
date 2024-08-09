package server

import (
	"context"
	"fmt"

	database "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driven"
	customerRepository "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driven/repositories/customer"
	itemRepository "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driven/repositories/item"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/config"
	customerHandler "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/http/customer"
	itemHandler "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/http/item"
	customerService "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/services/customer"
	itemService "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/services/item"
	"github.com/labstack/echo/v4"
)

func Start(cfg config.Config) {
	fmt.Println(context.Background(), fmt.Sprintf("Starting a server at http://%s", cfg.ServerHost))
	app := newApp(cfg)
	app.Logger.Fatal(app.Start(cfg.ServerHost))
}

func newApp(cfg config.Config) *echo.Echo {
	app := echo.New()

	database.ConectaDB(cfg.DatabaseConfig.Host, cfg.DatabaseConfig.User, cfg.DatabaseConfig.Password, cfg.DatabaseConfig.DbName, cfg.DatabaseConfig.Port)

	itemRepository := itemRepository.NewRepository(database.DB)
	itemService := itemService.NewItemService(itemRepository)
	itemHandler := itemHandler.NewItemHandler(itemService)
	itemHandler.RegisterRoutes(app)

	customerRepository := customerRepository.NewRepository(database.DB)
	customerService := customerService.NewCustomerService(customerRepository)
	customerHandler := customerHandler.NewCustomerHandler(customerService)
	customerHandler.RegisterRoutes(app)

	return app
}
