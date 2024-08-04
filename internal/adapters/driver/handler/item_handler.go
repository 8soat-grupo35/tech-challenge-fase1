package handler

import (
	"net/http"
	"strconv"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driven/repositories/item"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	services "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/services/item"
	"github.com/labstack/echo/v4"
)

//go:generate mockgen -source=item_handler.go -destination=../../../../test/adapters/driver/handler/mock/item_handler_mock.go
type ItemHandler interface {
	GetAll(echo echo.Context) error
	Create(echo echo.Context) error
	Update(echo echo.Context) error
	Delete(echo echo.Context) error
}

type itemHandler struct {
	*handler
}

func (h *handler) NewItemHandler() ItemHandler {
	return &itemHandler{h}
}

// GetAllItems godoc
// @Summary      List Items
// @Description  List All Items
// @Tags         items
// @Accept       json
// @Produce      json
// @Router       /items [get]
func (h *itemHandler) GetAll(echo echo.Context) error {
	var items []domain.Item

	category := echo.QueryParam("category")

	itemRepository := item.NewRepository(h.orm)
	service := services.NewItemService(itemRepository)

	items, err := service.GetAll(domain.Item{
		Category: category,
	})

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, items)
}

func (h *itemHandler) Create(echo echo.Context) error {
	itemDto := dto.ItemDto{}

	err := echo.Bind(&itemDto)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	itemRepository := item.NewRepository(h.orm)
	service := services.NewItemService(itemRepository)

	item, err := service.Create(domain.Item{
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

func (h *itemHandler) Update(echo echo.Context) error {
	itemDto := dto.ItemDto{}

	err := echo.Bind(&itemDto)
	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := strconv.Atoi(echo.Param("id"))

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	itemRepository := item.NewRepository(h.orm)
	service := services.NewItemService(itemRepository)

	item, err := service.Update(uint32(id), domain.Item{
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

func (h *itemHandler) Delete(echo echo.Context) error {
	id, err := strconv.Atoi(echo.Param("id"))

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	itemRepository := item.NewRepository(h.orm)
	service := services.NewItemService(itemRepository)

	err = service.Delete(uint32(id))

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, "item deleted successfully")
}
