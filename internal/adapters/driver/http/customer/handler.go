package customer

import (
	"net/http"
	"strconv"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	customerService "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/service"
	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	customerService customerService.CustomerService
}

func NewCustomerHandler(customerService customerService.CustomerService) CustomerHandler {
	return CustomerHandler{customerService: customerService}
}

func (c CustomerHandler) RegisterRoutes(server *echo.Echo) {
	customerGroupV1 := server.Group("/v1/customer")
	customerGroupV1.GET("", c.GetAll)
	customerGroupV1.POST("", c.Create)
	customerGroupV1.PUT("/:id", c.Update)
	customerGroupV1.GET("/cpf/:cpf", c.GetByCpf)
	customerGroupV1.DELETE("/:id", c.Delete)
}

func (h *CustomerHandler) GetAll(echo echo.Context) error {
	var customers []domain.Customer

	customers, err := h.customerService.GetAll()

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, customers)
}

// Create godoc
// @Summary      Insert Customer
// @Description  Insert Customer
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        CustomerToInsert	body dto.CustomerDto true "teste"
// @Router       /v1/customer [post]
// @success 200 {array} domain.Customer
// @Failure 500 {object} error
func (h *CustomerHandler) Create(echo echo.Context) error {
	customerDto := dto.CustomerDto{}

	err := echo.Bind(&customerDto)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	customer, err := h.customerService.Create(customerDto)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) Update(echo echo.Context) error {
	customerDto := dto.CustomerDto{}

	err := echo.Bind(&customerDto)
	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := strconv.Atoi(echo.Param("id"))

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	customer, err := h.customerService.Update(uint32(id), customerDto)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) Delete(echo echo.Context) error {
	id, err := strconv.Atoi(echo.Param("id"))

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.customerService.Delete(uint32(id))

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, "customer deleted successfully")
}

// GetByCpf godoc
// @Summary      Get Customer by CPF
// @Description  Retrieve a customer by their CPF
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        cpf   path      string  true  "CPF of the customer"
// @Router       /v1/customer/cpf/{cpf} [get]
// @Success      200  {object}  domain.Customer
// @Failure      500  {object}  error
func (h CustomerHandler) GetByCpf(echo echo.Context) error {

	cpf := echo.Param("cpf")

	customer, err := h.customerService.GetByCpf(cpf)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, customer)
}
