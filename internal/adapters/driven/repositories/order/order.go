package order

import (
	"log"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderRepository struct {
	orm *gorm.DB
}

func NewRepository(orm *gorm.DB) repository.OrderRepository {
	return &orderRepository{orm: orm}
}

func (c *orderRepository) GetAll() (orders []domain.Order, err error) {
	result := c.orm.Preload(clause.Associations).Preload("Items.Item").Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}}).Find(&orders)

	if result.Error != nil {
		log.Println(result.Error)
		return orders, result.Error
	}

	return orders, err
}

func (c *orderRepository) Create(order domain.Order) (*domain.Order, error) {
	result := c.orm.Create(&order)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &order, nil
}
