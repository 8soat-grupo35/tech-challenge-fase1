package domain_test

import (
	"testing"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewOrderSuccess(t *testing.T) {
	order := dto.OrderDto{
		Items: []dto.OrderItemDto{
			{Id: 1, Quantity: 2},
		},
		CustomerID: 1,
	}

	orderResult, err := domain.NewOrder(order)

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

	orderResult, err := domain.NewOrder(order)

	assert.Error(t, err)
	assert.EqualError(t, err, "customer_id: cannot be blank.")
	assert.Equal(t, orderResult, (*domain.Order)(nil))
}

func TestNewOrderWithoutItems(t *testing.T) {
	order := dto.OrderDto{
		Items:      []dto.OrderItemDto{},
		CustomerID: 1,
	}

	orderResult, err := domain.NewOrder(order)

	assert.Error(t, err)
	assert.EqualError(t, err, "items: cannot be blank.")
	assert.Equal(t, orderResult, (*domain.Order)(nil))
}
