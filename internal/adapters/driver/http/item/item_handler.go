package item

import (
	"net/http"
	"strconv"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	itemService "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/service"
	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	itemService itemService.ItemService
}

func NewItemHandler(itemService itemService.ItemService) ItemHandler {
	return ItemHandler{itemService: itemService}
}

func (h ItemHandler) RegisterRoutes(server *echo.Echo) {
	itemV1Group := server.Group("/v1/item")
	itemV1Group.GET("", h.GetAll)
	itemV1Group.POST("", h.Create)
	itemV1Group.PUT("/:id", h.Update)
	itemV1Group.DELETE("/:id", h.Delete)
}

func (h *ItemHandler) GetAll(echo echo.Context) error {
	var items []domain.Item

	category := echo.QueryParam("category")

	items, err := h.itemService.GetAll(domain.Item{
		Category: category,
	})

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, items)
}

func (h *ItemHandler) Create(echo echo.Context) error {
	itemDto := dto.ItemDto{}

	err := echo.Bind(&itemDto)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	item, err := h.itemService.Create(domain.Item{
		Name:     itemDto.Name,
		Category: itemDto.Category,
		Price:    itemDto.Price,
		ImageUrl: itemDto.ImageUrl,
	})

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, item)
}

func (h *ItemHandler) Update(echo echo.Context) error {
	itemDto := dto.ItemDto{}

	err := echo.Bind(&itemDto)
	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := strconv.Atoi(echo.Param("id"))

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	item, err := h.itemService.Update(uint32(id), domain.Item{
		Name:     itemDto.Name,
		Category: itemDto.Category,
		Price:    itemDto.Price,
		ImageUrl: itemDto.ImageUrl,
	})

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, item)
}

func (h *ItemHandler) Delete(echo echo.Context) error {
	id, err := strconv.Atoi(echo.Param("id"))

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.itemService.Delete(uint32(id))

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, "item deleted successfully")
}
