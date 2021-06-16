package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	SecKillTableName = "sec_kills"
	ProductTableName = "products"
	OrderTableName   = "orders"
)

type Product struct {
	gorm.Model
	ProductID     int    `gorm:"int:varchar(30);not null;comment:'产品id'"`
	ProductName   string `gorm:"type:varchar(100);not null;comment:'产品名'"`
	ProductNumber int    `gorm:"int:varchar(100);not null;comment:'产品数目'"`
}

type Order struct {
	gorm.Model
	OrderID   int       `gorm:"int:varchar(30);not null;comment:'orderid'"`
	OrderTime time.Time `json:"order_time" gorm:"column:order_time"`
	Payment   bool
	Product   Product `json:",omitempty" gorm:"foreignKey:ProductID"`
}

type SecKill struct {
	gorm.Model
	Product   Product   `json:",omitempty" gorm:"foreignKey:ProductID"`
	SecKillID int       `gorm:"int:varchar(30);not null;comment:'秒杀活动id'"`
	StartTime time.Time `json:"start_time" gorm:"column:start_time"`
	EndTime   time.Time `json:"end_time" gorm:"column:end_time"`
}

func ListSecKillJob() ([]SecKill, error) {
	var secKillList []SecKill
	res := Database.Table(SecKillTableName).Where("deleted_at is null ").Find(&secKillList)
	return secKillList, res.Error
}

func ListProducts() ([]Product, error) {
	var productList []Product
	res := Database.Table(ProductTableName).Where("deleted_at is null ").Find(&productList)
	return productList, res.Error
}

func ListOrder() ([]Order, error) {
	var orderList []Order
	res := Database.Table(OrderTableName).Where("deleted_at is null ").Find(&orderList)
	return orderList, res.Error
}
