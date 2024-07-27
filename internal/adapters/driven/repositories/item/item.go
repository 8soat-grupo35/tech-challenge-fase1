package item

import (
	"log"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports"
	"gorm.io/gorm"
)

type repository struct {
	orm *gorm.DB
}

func NewRepository(orm *gorm.DB) ports.ItemRepository {
	return &repository{orm: orm}
}

func (c *repository) GetAll() (items []domain.Item, err error) {
	result := c.orm.Find(&items)

	if result.Error != nil {
		log.Println(result.Error)
		return items, result.Error
	}

	return items, err
}

func (c *repository) GetOne(itemFilter domain.Item) (item *domain.Item, err error) {
	result := c.orm.Where(itemFilter).First(&item)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return item, nil
}

func (c *repository) Create(item domain.Item) (*domain.Item, error) {
	result := c.orm.Create(&item)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &item, nil
}

func (c *repository) Update(itemId uint32, item domain.Item) (*domain.Item, error) {
	itemModel := domain.Item{ID: itemId}
	result := c.orm.Model(&itemModel).Updates(&item)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &itemModel, nil
}

func (c *repository) Delete(itemId uint32) error {
	result := c.orm.Delete(&domain.Item{}, itemId)

	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}
