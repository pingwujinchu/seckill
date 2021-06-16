package api

import (
	"encoding/json"
	"log"
	"server/pkg/cache"
	"server/pkg/controller/dto"
	models "server/pkg/model"
	RabbitMQ "server/pkg/rabitmq"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/gofrs/uuid"
)

const (
	SecSkillQueue = "seckill"
)

var secKillRabbitmq = RabbitMQ.NewRabbitMQSimple(SecSkillQueue)

//products
func Products(c *gin.Context) {
	var productList []models.Product
	val, err := cache.Rdb.Get("products").Result()
	if err != nil {
		productList, err = models.ListProducts()
		log.Println("List product Test Success.")
		productJson, _ := json.Marshal(productList)
		cache.Rdb.Set("products", productJson, time.Hour)
	} else {
		json.Unmarshal([]byte(val), &productList)
	}

	dto.APIResponse(c, err, productList)
}

//seckills
func Seckills(c *gin.Context) {
	secKillList, err := models.ListSecKillJob()
	log.Println("List seckills Test Success.")
	dto.APIResponse(c, err, secKillList)
}

//Order
func Orders(c *gin.Context) {
	orderList, err := models.ListOrder()
	log.Println("List order Test Success.")
	dto.APIResponse(c, err, orderList)
}

// 秒杀前先检查缓存中是否有库存
func SecKillProduct(c *gin.Context) {
	request_id, _ := uuid.NewV4()
	var productList []models.Product
	val, err := cache.Rdb.Get("products").Result()
	if err != nil {
		productList, err = models.ListProducts()
		log.Println("List product Test Success.")
		productJson, _ := json.Marshal(productList)
		cache.Rdb.Set("products", productJson, time.Hour)
	} else {
		json.Unmarshal([]byte(val), &productList)
	}

	productId := c.Param("productId")
	id, err := strconv.Atoi(productId)
	var currProduct models.Product
	for _, product := range productList {
		if product.ProductID == id {
			currProduct = product
		}
	}

	if currProduct.ProductNumber <= 0 {
		e := errors.New(-1, "商品已经卖完")
		dto.APIResponse(c, e, "商品已经卖完")
	}
	secKillRabbitmq.PublishSimple(request_id.String())
	dto.APIResponse(c, err, request_id.String()+"正在排队中")
}

// 秒杀后进行支付
func payForProduct(c *gin.Context) {
	var productList []models.Product
	val, err := cache.Rdb.Get("products").Result()
	if err != nil {
		productList, err = models.ListProducts()
		log.Println("List product Test Success.")
		productJson, _ := json.Marshal(productList)
		cache.Rdb.Set("products", productJson, time.Hour)
	} else {
		json.Unmarshal([]byte(val), &productList)
	}

	productId := c.Param("productId")
	id, err := strconv.Atoi(productId)
	var currProduct models.Product
	for _, product := range productList {
		if product.ProductID == id {
			currProduct = product
		}
	}

	if currProduct.ProductNumber <= 0 {
		e := errors.New(-1, "商品已经卖完")
		dto.APIResponse(c, e, "商品已经卖完")
	}
	dto.APIResponse(c, err, productList)

}

//获取当前状态
func GetStatus(c *gin.Context) {
	var status int
	requestID := c.Param("requestId")
	val, err := cache.Rdb.Get("status/" + requestID).Result()
	if err != nil {
		cache.Rdb.Set("status/"+requestID, "0", time.Hour)
		status = 0
	} else {
		status, _ = strconv.Atoi(val)
	}

	switch status {
	case 0:
		dto.APIResponse(c, nil, "请求已提交")
		break
	case 1:
		dto.APIResponse(c, nil, "秒杀排队中")
		break
	case 2:
		dto.APIResponse(c, nil, "已锁定，等待支付")
		break
	case 3:
		dto.APIResponse(c, nil, "已支付")
		break
	case -1:
		dto.APIResponse(c, nil, "没抢到")
		break
	default:
		dto.APIResponse(c, nil, "没抢到")
	}
}

//获取当前状态
func GetAllStatus(c *gin.Context) {

	// val, err := cache.Rdb.Get("status").Result()
	// if err != nil {
	// 	productList, err = models.ListProducts()
	// 	log.Println("List product Test Success.")
	// 	productJson, _ := json.Marshal(productList)
	// 	cache.Rdb.Set("products", productJson, time.Hour)
	// } else {
	// 	json.Unmarshal([]byte(val), &productList)
	// }
}
