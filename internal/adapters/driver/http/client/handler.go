package client

import (
	"net/http"
	"strconv"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	clientService "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/service"
	"github.com/labstack/echo/v4"
)

type ClientHandler struct {
	clientService clientService.ClientService
}

func NewClientHandler(clientService clientService.ClientService) ClientHandler {
	return ClientHandler{clientService: clientService}
}

func (c ClientHandler) RegisterRoutes(server *echo.Echo) {
	clientGroupV1 := server.Group("/v1/client")
	clientGroupV1.GET("", c.GetAll)
	clientGroupV1.POST("", c.Create)
	clientGroupV1.GET("/cpf/:cpf", c.GetByCpf)
	clientGroupV1.DELETE("/:id", c.Delete)
}

func (h *ClientHandler) GetAll(echo echo.Context) error {
	var clients []domain.Client

	clients, err := h.clientService.GetAll()

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, clients)
}

func (h *ClientHandler) Create(echo echo.Context) error {
	clientDto := dto.ClientDto{}

	err := echo.Bind(&clientDto)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	client, err := h.clientService.Create(domain.Client{
		Name:  clientDto.Name,
		Email: clientDto.Email,
		CPF:   clientDto.CPF,
	})

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, client)
}

func (h *ClientHandler) Update(echo echo.Context) error {
	clientDto := dto.ClientDto{}

	err := echo.Bind(&clientDto)
	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := strconv.Atoi(echo.Param("id"))

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	client, err := h.clientService.Update(uint32(id), domain.Client{
		Name:  clientDto.Name,
		Email: clientDto.Email,
		CPF:   clientDto.CPF,
	})

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, client)
}

func (h *ClientHandler) Delete(echo echo.Context) error {
	id, err := strconv.Atoi(echo.Param("id"))

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.clientService.Delete(uint32(id))

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, "client deleted successfully")
}

func (h ClientHandler) GetByCpf(echo echo.Context) error {

	cpf := echo.Param("cpf")

	client, err := h.clientService.GetByCpf(cpf)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, client)
}
