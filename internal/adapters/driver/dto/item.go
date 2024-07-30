package dto

type ItemDto struct {
	Id       uint32  `param:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float32 `json:"price"`
	ImageUrl string  `json:"image_url"`
}
