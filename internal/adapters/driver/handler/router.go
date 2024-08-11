package handler

import (
	_ "github.com/8soat-grupo35/tech-challenge-fase1/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *echo.Echo {
	r := echo.New()

	mapRoutes(r, db)

	return r
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
// @BasePath /
func mapRoutes(echo *echo.Echo, orm *gorm.DB) {
	handler := NewHandler(orm)

	echo.GET("/swagger/*", echoSwagger.WrapHandler)

	echo.GET("/items", handler.NewItemHandler().GetAll)
	echo.POST("/item", handler.NewItemHandler().Create)
	echo.PUT("/item/:id", handler.NewItemHandler().Update)
	echo.DELETE("/item/:id", handler.NewItemHandler().Delete)

	echo.GET("/orders", handler.NewOrderHandler().GetAll)
	echo.POST("/checkout", handler.NewOrderHandler().Checkout)

	echo.Logger.Fatal(echo.Start(":8000"))
}
