package handler

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

// GetAll godoc
// @Summary      List Items
// @Description  List All Items
// @Tags         Items
// @Accept       json
// @Produce      json
// @Router       /items [get]
// @success 200  {object} domain.Item
// @Failure 500 {object} error
func (h *ItemHandler) GetAll(echo echo.Context) error {
	var items []domain.Item

	category := echo.QueryParam("category")

	items, err := h.itemService.GetAll(category)

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
func (h *ItemHandler) Create(echo echo.Context) error {
	itemDto := dto.ItemDto{}

	err := echo.Bind(&itemDto)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	item, err := h.itemService.Create(itemDto)

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

	item, err := h.itemService.Update(uint32(id), itemDto)

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
