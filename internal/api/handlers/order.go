package handlers

import (
	"net/http"
	"strconv"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/dto"
	custom_errors "github.com/8soat-grupo35/tech-challenge-fase1/internal/api/errors"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/controllers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
// @success 200 {array} presenters.OrderPresenter
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

// Create godoc
// @Summary      Update Order Status
// @Description  Update Order Status
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param		 		 id     path int         true "ID do item"
// @Param        Order	body dto.OrderStatusDto true "Status to update Order"
// @Router       /v1/orders/{id} [patch]
// @success 200 {object} domain.Order
// @Failure 500 {object} error
func (h *OrderHandler) UpdateStatus(echo echo.Context) error {
	orderDto := dto.OrderDto{}

	id, err := strconv.Atoi(echo.Param("id"))
	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	bindError := echo.Bind(&orderDto)
	if bindError != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	order, err := h.orderController.UpdateStatus(uint32(id), orderDto.Status)
	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	if order == nil {
		return echo.JSON(http.StatusNoContent, err.Error())
	}

	return echo.JSON(http.StatusOK, order)
}

// GetOrderPaymentStatus godoc
// @Summary      Get Order Payment Status
// @Description  Get Order Payment Status
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param		 orderID             path int         true "ID do pedido"
// @Router       /v1/orders/{orderID}/payment/status [get]
// @success 200 {object} presenters.OrderPaymentStatusPresenter
// @Failure 500 {object} error
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

// UpdateOrderPaymentStatus godoc
// @Summary      Update Order Payment Status
// @Description  Update Order Payment Status
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param		 orderID             path int         true "ID do pedido"
// @Param        UpdateBody	body dto.OrderPaymentStatusDto true "UpdateBody"
// @Router       /v1/orders/{orderID}/payment/status [put]
// @success 200 {object} string
// @Failure 500 {object} error
func (h *OrderHandler) UpdateOrderPaymentStatus(echo echo.Context) error {
	paramOrderID := echo.Param("orderID")
	orderID, err := strconv.Atoi(paramOrderID)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, &custom_errors.BadRequestError{
			Message: "Order ID was not integer",
		})
	}

	body := dto.OrderPaymentStatusDto{}

	err = echo.Bind(&body)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, &custom_errors.BadRequestError{
			Message: "Body could not be parsed",
		})
	}

	err = h.orderController.UpdatePaymentStatus(uint32(orderID), body.Status)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, "Payment Status Updated!")
}
