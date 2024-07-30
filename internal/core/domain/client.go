package domain

import "gorm.io/gorm"

type Client struct {
	Id    uint32 `gorm:"primary_key;auto_increment"`
	Name  string `gorm:"size:255;not null;"`
	Email string `gorm:"size:255;not null;"`
	CPF   string `gorm:"size:11;not null;"`
	gorm.Model
}
