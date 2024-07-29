package handler

import (
	"gorm.io/gorm"
)

//go:generate mockgen -source=handler.go -destination=../../../../test/adapters/driver/handler/mock/handler_mock.go
type Handlers interface {
	NewItemHandler() ItemHandler
}

type handler struct {
	orm *gorm.DB
}

func NewHandler(orm *gorm.DB) Handlers {
	return &handler{orm}
}
