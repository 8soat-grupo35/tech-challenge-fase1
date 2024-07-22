package http

import (
	"net/http"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports"
	"github.com/labstack/echo"
)

type Handler struct {
	service ports.ItemService
}

func NewHandler(service ports.ItemService, echo *echo.Echo) *Handler {
	handler := Handler{service: service}

	echo.GET("/items", handler.GetAll)

	return &handler
}

func (h *Handler) GetAll(echo echo.Context) error {
	var items []domain.Item

	items, err := h.service.GetAll()

	if err != nil {
		echo.Error(err)
	}

	return echo.JSON(http.StatusOK, items)
}
