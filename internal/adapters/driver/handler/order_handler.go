package handler

import (
	"net/http"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driven/repositories/order"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/dto"
	services "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/services/order"
	"github.com/labstack/echo/v4"
)

//go:generate mockgen -source=order_handler.go -destination=../../../../test/adapters/driver/handler/mock/order_handler_mock.go
type OrderHandler interface {
	GetAll(echo echo.Context) error
	Checkout(echo echo.Context) error
}

type orderHandler struct {
	*handler
}

func (h *handler) NewOrderHandler() OrderHandler {
	return &orderHandler{h}
}

// GetAll godoc
// @Summary      List Orders
// @Description  List All Orders
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Router       /orders [get]
// @Success 200  {object} domain.Order
// @Failure 500  {object} error
func (h *orderHandler) GetAll(echo echo.Context) error {
	orderRepository := order.NewRepository(h.orm)
	service := services.NewOrderService(orderRepository)

	orders, err := service.GetAll()

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
// @Router       /checkout [post]
// @success 200 {array} domain.Order
// @Failure 500 {object} error
func (h *orderHandler) Checkout(echo echo.Context) error {
	orderDto := dto.OrderDto{}

	err := echo.Bind(&orderDto)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	orderRepository := order.NewRepository(h.orm)
	service := services.NewOrderService(orderRepository)

	order, err := service.Create(orderDto)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, order)
}
