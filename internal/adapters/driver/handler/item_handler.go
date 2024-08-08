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

// GetAll godoc
// @Summary      List Items
// @Description  List All Items
// @Tags         Items
// @Accept       json
// @Produce      json
// @Router       /items [get]
// @success 200  {object} domain.Item
// @Failure 500 {object} error
func (h *itemHandler) GetAll(echo echo.Context) error {
	var items []domain.Item

	category := echo.QueryParam("category")

	itemRepository := item.NewRepository(h.orm)
	service := services.NewItemService(itemRepository)

	items, err := service.GetAll(category)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, items)
}

// Create godoc
// @Summary      Insert Item
// @Description  Insert Item
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        ItemToInsert	body dto.ItemDto true "teste"
// @Router       /item [post]
// @success 200 {array} domain.Item
// @Failure 500 {object} error
func (h *itemHandler) Create(echo echo.Context) error {
	itemDto := dto.ItemDto{}

	err := echo.Bind(&itemDto)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	itemRepository := item.NewRepository(h.orm)
	service := services.NewItemService(itemRepository)

	item, err := service.Create(itemDto)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, item)
}

// Update godoc
// @Summary      Update Item
// @Description  Update Item
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param		 Id             path int         true "ID do item"
// @Param        ItemToInsert	body dto.ItemDto true "teste"
// @Router       /item/{id} [put]
// @success 200 {array} domain.Item
// @Failure 500 {object} error
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

	item, err := service.Update(uint32(id), itemDto)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, item)
}

// Delete godoc
// @Summary      Delete Item
// @Description  Delete Item
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param		 Id             path int         true "ID do item"
// @Router       /item/{id} [delete]
// @success 200 {string}  string    "item deleted successfully"
// @Failure 500 {object} error
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
