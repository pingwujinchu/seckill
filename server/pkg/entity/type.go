package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductID     int    `gorm:"int:varchar(30);not null;comment:'产品id'"`
	ProductName   string `gorm:"type:varchar(100);not null;comment:'产品名'"`
	ProductNumber int    `gorm:"int:varchar(100);not null;comment:'产品数目'"`
}

type Order struct {
	gorm.Model
	OrderID   int
	OrderDate date
}
