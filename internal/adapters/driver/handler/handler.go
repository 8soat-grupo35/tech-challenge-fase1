package handler

import (
	"gorm.io/gorm"
)

type Handlers interface {
	Item() ItemHandler
}

type handler struct {
	orm *gorm.DB
}

func NewHandler(orm *gorm.DB) Handlers {
	return &handler{orm}
}
