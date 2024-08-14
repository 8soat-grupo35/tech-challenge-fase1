package handler

import (
	"net/http"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/dto"
	service "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/service"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) OrderHandler {
	return OrderHandler{orderService: orderService}
}

func (h OrderHandler) RegisterRoutes(server *echo.Echo) {
	orderV1Group := server.Group("/v1/orders")
	orderV1Group.GET("", h.GetAll)
	orderV1Group.POST("/checkout", h.Checkout)
}

// GetAll godoc
// @Summary      List Orders
// @Description  List All Orders
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Router       /v1/orders [get]
// @Success 200  {object} domain.Order
// @Failure 500  {object} error
func (h *OrderHandler) GetAll(echo echo.Context) error {
	orders, err := h.orderService.GetAll()

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, orders)
}

// Create godoc
// @Summary      Insert Order
// @Description  Insert Order
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        Order	body dto.OrderDto true "Order to create"
// @Router       /v1/orders/checkout [post]
// @success 200 {array} domain.Order
// @Failure 500 {object} error
func (h *OrderHandler) Checkout(echo echo.Context) error {
	orderDto := dto.OrderDto{}

	err := echo.Bind(&orderDto)
	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	order, err := h.orderService.Create(orderDto)
	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, order)
}
