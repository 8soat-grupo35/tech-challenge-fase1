package gateways

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/src/entities"
	"github.com/8soat-grupo35/tech-challenge-fase1/src/interfaces/repository"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderGateway struct {
	orm *gorm.DB
}

func NewOrderGateway(orm *gorm.DB) repository.OrderRepository {
	return &orderGateway{orm: orm}
}

func (c *orderGateway) GetAll() (orders []entities.Order, err error) {
	result := c.orm.Preload(clause.Associations).Preload("Items.Item").Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}}).Find(&orders)

	if result.Error != nil {
		log.Println(result.Error)
		return orders, result.Error
	}

	return orders, err
}

func (c *orderGateway) Create(order entities.Order) (*entities.Order, error) {
	result := c.orm.Create(&order)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &order, nil
}
