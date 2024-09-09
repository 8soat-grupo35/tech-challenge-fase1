package entities

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/src/adapters/dto"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type OrderItem struct {
	ID       uint `gorm:"primarykey;autoIncrement" json:"-"`
	OrderID  uint `json:"-"`
	ItemID   uint `json:"id"`
	Quantity uint `json:"quantity"`
	Item     Item `gorm:"references:ID" json:"-"`
} //@name domain.OrderItem

type Order struct {
	ID         uint        `gorm:"primarykey;autoIncrement" json:"id"`
	Items      []OrderItem `gorm:"foreignKey:OrderID;references:ID;constraint:OnDelete:CASCADE" json:"items"`
	CustomerID uint32      `json:"customer_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} //@name domain.Order

func NewOrder(orderDto dto.OrderDto) (*Order, error) {
	newOrder := Order{
		CustomerID: orderDto.CustomerID,
		Items:      OrderItemToDomain(orderDto),
	}

	err := newOrder.Validate()

	if err != nil {
		return nil, err
	}

	return &newOrder, err
}

func OrderItemToDomain(orderDto dto.OrderDto) (list []OrderItem) {

	for _, orderItemDto := range orderDto.Items {
		list = append(list, OrderItem{
			ItemID:   orderItemDto.Id,
			Quantity: orderItemDto.Quantity,
		})
	}

	return list
}

func (order Order) Validate() error {
	return validation.ValidateStruct(
		&order,
		validation.Field(
			&order.Items,
			validation.Required.When(len(order.Items) > 0).Error("must be one or more item"),
			validation.Required,
		),
		validation.Field(
			&order.CustomerID,
			validation.Required,
		),
	)
}
