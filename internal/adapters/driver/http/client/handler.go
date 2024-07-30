package client

import "github.com/labstack/echo/v4"

type ClientHandler struct {
}

func (c ClientHandler) RegisterRoutes(server *echo.Echo) {
	clientGroupV1 := server.Group("/v1/client")
	clientGroupV1.GET("", c.GetAll)
	clientGroupV1.POST("", c.Create)
	clientGroupV1.GET("/:id", c.GetById)
	clientGroupV1.GET("/cpf/:cpf", c.GetByCpf)
	clientGroupV1.DELETE("/:id", c.Delete)
}

func (c ClientHandler) GetAll(echoContext echo.Context) error {
	return nil
}

func (c ClientHandler) Create(echoContext echo.Context) error {
	return nil
}

func (c ClientHandler) GetById(echoContext echo.Context) error {
	return nil
}

func (c ClientHandler) GetByCpf(echoContext echo.Context) error {
	return nil
}

func (c ClientHandler) Delete(echoContext echo.Context) error {
	return nil
}
