package server

import (
	"context"
	"fmt"

	_ "github.com/8soat-grupo35/tech-challenge-fase1/docs"
	database "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driven"
	customerRepository "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driven/repositories/customer"
	itemRepository "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driven/repositories/item"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/config"
	customerHandler "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/http/customer"
	itemHandler "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/http/item"
	customerService "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/services/customer"
	itemService "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/services/item"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Start(cfg config.Config) {
	fmt.Println(context.Background(), fmt.Sprintf("Starting a server at http://%s", cfg.ServerHost))
	app := newApp(cfg)
	app.Logger.Fatal(app.Start(cfg.ServerHost))
}

// @title Swagger Fastfood App API
// @version 1.0
// @description This is a sample API from Fastfood App.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /v1
func newApp(cfg config.Config) *echo.Echo {
	database.ConectaDB(cfg.DatabaseConfig.Host, cfg.DatabaseConfig.User, cfg.DatabaseConfig.Password, cfg.DatabaseConfig.DbName, cfg.DatabaseConfig.Port)

	app := echo.New()
	app.GET("/swagger/*", echoSwagger.WrapHandler)

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
