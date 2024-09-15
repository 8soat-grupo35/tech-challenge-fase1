package handlers

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/dto"
	custom_errors "github.com/8soat-grupo35/tech-challenge-fase1/internal/api/errors"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/controllers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func (h *OrderHandler) GetOrderPaymentStatus(echo echo.Context) error {
	paramOrderID := echo.Param("orderID")
	orderID, err := strconv.Atoi(paramOrderID)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, &custom_errors.BadRequestError{
			Message: "Order ID was not integer",
		})
	}

	presenter, err := h.orderController.GetPaymentStatus(uint32(orderID))

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, presenter)
}
