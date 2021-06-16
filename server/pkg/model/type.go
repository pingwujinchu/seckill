package models

import (
	"encoding/json"
	"log"
	"server/pkg/cache"
	"time"

	"gorm.io/gorm"
)

const (
	SecKillTableName = "sec_kills"
	ProductTableName = "products"
	OrderTableName   = "orders"
)

type Message struct {
	ProductID int
	RequestID string
}

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
	RequestID string
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

//使用事务操作，先减少库存，然后再更新缓存
func SolveSecKill(requestID string, ProductID int) {
	tx := Database.Begin()
	var product Product
	tx.Where("ProductID=" + string(ProductID)).First(&product)
	if product.ProductNumber > 0 {
		product.ProductNumber = product.ProductNumber - 1
	}
	tx.Save(product)
	order := Order{
		OrderTime: time.Now(),
		RequestID: requestID,
		Product:   product,
	}
	tx.Save(order)

	var productList []Product
	val, err := cache.Rdb.Get("products").Result()
	if err != nil {
		productList, err = ListProducts()
		log.Println("List product Test Success.")
		productJson, _ := json.Marshal(productList)
		cache.Rdb.Set("products", productJson, time.Hour)
	} else {
		json.Unmarshal([]byte(val), &productList)
		for _, p := range productList {
			if p.ProductID == ProductID {
				p.ProductNumber = p.ProductNumber - 1
			}
		}
		productJson, _ := json.Marshal(productList)
		cache.Rdb.Set("products", productJson, time.Hour)
	}
	cache.Rdb.Set("status/"+requestID, "2", time.Hour)
	tx.Commit()
}
