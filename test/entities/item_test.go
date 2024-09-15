package entities_test

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewItemSuccess(t *testing.T) {

	itemTest := dto.ItemDto{
		Name:     "Milkshake",
		Category: "SOBREMESA",
		Price:    20,
		ImageUrl: "https://blog.biglar.com.br/wp-content/uploads/2022/08/iStock-1308045723.jpeg",
	}

	itemResult, err := entities.NewItem(itemTest)

	assert.NoError(t, err)
	assert.Equal(t, itemTest.Name, itemResult.Name)
	assert.Equal(t, itemTest.Category, itemResult.Category)
	assert.Equal(t, itemTest.Price, itemResult.Price)
	assert.Equal(t, itemTest.ImageUrl, itemResult.ImageUrl)
}

func TestNewItemNameNotFound(t *testing.T) {
	itemTest := dto.ItemDto{
		Name:     "",
		Category: "SOBREMESA",
		Price:    20,
		ImageUrl: "https://blog.biglar.com.br/wp-content/uploads/2022/08/iStock-1308045723.jpeg",
	}

	itemResult, err := entities.NewItem(itemTest)

	assert.Error(t, err)
	assert.EqualError(t, err, "Name: cannot be blank.")
	assert.Equal(t, itemResult, (*entities.Item)(nil))
}

func TestNewItemNameInvalidLenght(t *testing.T) {
	itemTest := dto.ItemDto{
		Name:     "AA",
		Category: "SOBREMESA",
		Price:    20,
		ImageUrl: "https://blog.biglar.com.br/wp-content/uploads/2022/08/iStock-1308045723.jpeg",
	}

	itemResult, err := entities.NewItem(itemTest)

	assert.Error(t, err)
	assert.EqualError(t, err, "Name: the length must be between 3 and 255.")
	assert.Equal(t, itemResult, (*entities.Item)(nil))
}

func TestNewItemCategoryNotFound(t *testing.T) {
	itemTest := dto.ItemDto{
		Name:     "Milkshake",
		Price:    20,
		ImageUrl: "https://blog.biglar.com.br/wp-content/uploads/2022/08/iStock-1308045723.jpeg",
	}

	itemResult, err := entities.NewItem(itemTest)

	assert.Error(t, err)
	assert.EqualError(t, err, "Category: cannot be blank.")
	assert.Equal(t, itemResult, (*entities.Item)(nil))
}

func TestNewItemCategoryNotInAllowedValues(t *testing.T) {
	itemTest := dto.ItemDto{
		Name:     "Milkshake",
		Category: "APERITIVO",
		Price:    20,
		ImageUrl: "https://blog.biglar.com.br/wp-content/uploads/2022/08/iStock-1308045723.jpeg",
	}

	itemResult, err := entities.NewItem(itemTest)

	assert.Error(t, err)
	assert.EqualError(t, err, "Category: must be a valid value between (lanche,sobremesa,acompanhamento,bebida).")
	assert.Equal(t, itemResult, (*entities.Item)(nil))
}

func TestNewItemPriceNotFound(t *testing.T) {
	itemTest := dto.ItemDto{
		Name:     "Milkshake",
		Category: "SOBREMESA",
		ImageUrl: "https://blog.biglar.com.br/wp-content/uploads/2022/08/iStock-1308045723.jpeg",
	}

	itemResult, err := entities.NewItem(itemTest)

	assert.Error(t, err)
	assert.EqualError(t, err, "Price: cannot be blank.")
	assert.Equal(t, itemResult, (*entities.Item)(nil))
}

func TestNewItemPriceBelowMinimunValue(t *testing.T) {
	itemTest := dto.ItemDto{
		Name:     "Milkshake",
		Category: "SOBREMESA",
		Price:    -2,
		ImageUrl: "https://blog.biglar.com.br/wp-content/uploads/2022/08/iStock-1308045723.jpeg",
	}

	itemResult, err := entities.NewItem(itemTest)

	assert.Error(t, err)
	assert.EqualError(t, err, "Price: must be no less than 0.01.")
	assert.Equal(t, itemResult, (*entities.Item)(nil))
}

func TestNewItemImageNotFound(t *testing.T) {
	itemTest := dto.ItemDto{
		Name:     "Milkshake",
		Category: "SOBREMESA",
		Price:    20,
	}

	itemResult, err := entities.NewItem(itemTest)

	assert.Error(t, err)
	assert.EqualError(t, err, "ImageUrl: cannot be blank.")
	assert.Equal(t, itemResult, (*entities.Item)(nil))
}

func TestNewItemImageNotInAURLFormat(t *testing.T) {
	itemTest := dto.ItemDto{
		Name:     "Milkshake",
		Category: "SOBREMESA",
		Price:    20,
		ImageUrl: "asiodjaoiwdjioasjdiousjs9",
	}

	itemResult, err := entities.NewItem(itemTest)

	assert.Error(t, err)
	assert.EqualError(t, err, "ImageUrl: must be a valid URL.")
	assert.Equal(t, itemResult, (*entities.Item)(nil))
}
