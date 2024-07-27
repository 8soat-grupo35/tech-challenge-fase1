package domain

import "gorm.io/gorm"

type Item struct {
	ID       uint32  `gorm:"primary_key;auto_increment"`
	Name     string  `gorm:"size:255;not null;"`
	Category string  `gorm:"size:30;not null;"`
	Price    float32 `gorm:"not null;"`
	ImageUrl string  `gorm:"size:255;not null;"`
	gorm.Model
}
