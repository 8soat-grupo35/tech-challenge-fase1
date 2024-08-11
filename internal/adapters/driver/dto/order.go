package dto

type OrderItemDto struct {
	Id       uint `json:"id"`
	Quantity uint `json:"quantity"`
} //@name OrderItemDto

type OrderDto struct {
	Items      []OrderItemDto `json:"items"`
	CustomerID uint32         `json:"customer_id"`
} //@name OrderDto
