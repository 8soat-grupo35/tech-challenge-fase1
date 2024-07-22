package domain

import "gorm.io/gorm"

type Item struct {
	Id   uint32 `gorm:"primary_key;auto_increment"`
	Name string `gorm:"size:255;not null;"`
	gorm.Model
}
