package entities_test

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/src/adapters/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/src/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrderSuccess(t *testing.T) {
	order := dto.OrderDto{
		Items: []dto.OrderItemDto{
			{Id: 1, Quantity: 2},
		},
		CustomerID: 1,
	}

	orderResult, err := entities.NewOrder(order)

	assert.NoError(t, err)
	assert.Equal(t, len(order.Items), len(orderResult.Items))
	assert.Equal(t, order.CustomerID, orderResult.CustomerID)
}

func TestNewOrderWithoutCustomer(t *testing.T) {
	order := dto.OrderDto{
		Items: []dto.OrderItemDto{
			{Id: 1, Quantity: 2},
		},
	}

	orderResult, err := entities.NewOrder(order)

	assert.Error(t, err)
	assert.EqualError(t, err, "customer_id: cannot be blank.")
	assert.Equal(t, orderResult, (*entities.Order)(nil))
}

func TestNewOrderWithoutItems(t *testing.T) {
	order := dto.OrderDto{
		Items:      []dto.OrderItemDto{},
		CustomerID: 1,
	}

	orderResult, err := entities.NewOrder(order)

	assert.Error(t, err)
	assert.EqualError(t, err, "items: cannot be blank.")
	assert.Equal(t, orderResult, (*entities.Order)(nil))
}
