package handlers

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/src/adapters/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/src/controllers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type OrderHandler struct {
	orderController *controllers.OrderController
}

func NewOrderHandler(db *gorm.DB) OrderHandler {
	return OrderHandler{
		orderController: controllers.NewOrderController(db),
	}
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
	orders, err := h.orderController.GetAll()

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

	order, err := h.orderController.Checkout(orderDto)
	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, order)
}