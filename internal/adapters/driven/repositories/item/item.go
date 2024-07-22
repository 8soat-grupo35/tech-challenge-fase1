package item

import (
	"log"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	"gorm.io/gorm"
)

type Repository struct {
	orm *gorm.DB
}

func NewRepository(orm *gorm.DB) *Repository {
	return &Repository{orm: orm}
}

func (c *Repository) GetAll() (items []domain.Item, err error) {
	result := c.orm.Find(&items)

	if result.Error != nil {
		log.Println(result.Error)
		return items, result.Error
	}

	return items, err
}
