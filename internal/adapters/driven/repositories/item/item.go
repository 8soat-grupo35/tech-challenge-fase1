package item

import (
	"log"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/repository"
	"gorm.io/gorm"
)

type itemRepository struct {
	orm *gorm.DB
}

func NewRepository(orm *gorm.DB) repository.ItemRepository {
	return &itemRepository{orm: orm}
}

func (c *itemRepository) GetAll(filter domain.Item) (items []domain.Item, err error) {
	result := c.orm.Where(filter).Find(&items)

	if result.Error != nil {
		log.Println(result.Error)
		return items, result.Error
	}

	return items, err
}

func (c *itemRepository) GetOne(itemFilter domain.Item) (item *domain.Item, err error) {
	result := c.orm.Where(itemFilter).First(&item)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return item, nil
}

func (c *itemRepository) Create(item domain.Item) (*domain.Item, error) {
	result := c.orm.Create(&item)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &item, nil
}

func (c *itemRepository) Update(itemId uint32, item domain.Item) (*domain.Item, error) {
	itemModel := domain.Item{ID: itemId}
	result := c.orm.Model(&itemModel).Updates(&item)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &itemModel, nil
}

func (c *itemRepository) Delete(itemId uint32) error {
	result := c.orm.Delete(&domain.Item{}, itemId)

	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}
