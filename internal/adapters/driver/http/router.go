package http

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driven/repositories/item"
	itemHandler "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/http/item"
	itemService "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/services/item"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *echo.Echo {
	r := echo.New()

	mapRoutes(r, db)

	return r
}

func mapRoutes(echo *echo.Echo, orm *gorm.DB) {

	// Injections
	// Repositories
	itemRepository := item.NewRepository(orm)

	// Services
	itemService := itemService.New(itemRepository)

	// Handlers
	itemHandler.NewHandler(itemService, echo)

	echo.Logger.Fatal(echo.Start(":8000"))
}
