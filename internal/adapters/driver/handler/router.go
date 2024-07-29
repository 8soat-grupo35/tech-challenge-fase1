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

	echo.GET("/items", handler.NewItemHandler().GetAll)
	echo.POST("/item", handler.NewItemHandler().Create)
	echo.PUT("/item/:id", handler.NewItemHandler().Update)
	echo.DELETE("/item/:id", handler.NewItemHandler().Delete)

	echo.Logger.Fatal(echo.Start(":8000"))
}
