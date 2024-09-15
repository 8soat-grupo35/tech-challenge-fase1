package dto

type OrderItemDto struct {
	Id       uint32 `json:"id"`
	Quantity uint32 `json:"quantity"`
} //@name OrderItemDto

type OrderDto struct {
	Items      []OrderItemDto `json:"items"`
	CustomerID uint32         `json:"customer_id"`
} //@name OrderDto
