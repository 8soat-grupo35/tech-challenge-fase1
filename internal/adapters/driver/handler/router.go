package handler

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *echo.Echo {
	r := echo.New()

	mapRoutes(r, db)

	return r
}

func mapRoutes(echo *echo.Echo, orm *gorm.DB) {
	handler := NewHandler(orm)

	echo.GET("/items", handler.Item().GetAll)
	echo.POST("/item", handler.Item().Create)
	echo.PUT("/item/:id", handler.Item().Update)
	echo.DELETE("/item/:id", handler.Item().Delete)

	echo.Logger.Fatal(echo.Start(":8000"))
}
