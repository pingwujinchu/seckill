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
	var secKillList []models.SecKill
	val, err := cache.Rdb.Get("seckills").Result()
	if err != nil {
		secKillList, err = models.ListSecKillJob()
		log.Println("List seckills Test Success.")
		seckillJson, _ := json.Marshal(secKillList)
		cache.Rdb.Set("seckills", seckillJson, time.Hour)
	} else {
		json.Unmarshal([]byte(val), &secKillList)
	}
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
	log.Println("pid:", productId)
	id, err := strconv.Atoi(productId)
	var currProduct models.Product

	var secKillList []models.SecKill
	val, err = cache.Rdb.Get("seckills").Result()
	if err != nil {
		secKillList, err = models.ListSecKillJob()
		log.Println("List seckills Test Success.")
		seckillJson, _ := json.Marshal(secKillList)
		cache.Rdb.Set("seckills", seckillJson, time.Hour)
	} else {
		json.Unmarshal([]byte(val), &secKillList)
	}

	if !checkTime(id, secKillList) {
		dto.APIResponse(c, nil, "不在秒杀时间内")
		return
	}

	for _, product := range productList {
		log.Println(product)
		if int(product.ID) == id {
			currProduct = product
		}
	}

	log.Println("name: ", currProduct.ProductName)
	log.Println("number: ", currProduct.ProductNumber)
	if currProduct.ProductNumber <= 0 {
		e := errors.New(-1, "商品已经卖完")
		dto.APIResponse(c, e, "商品已经卖完")
		return
	}

	message := models.Message{
		ProductID: int(currProduct.ID),
		RequestID: request_id.String(),
	}
	jsonMessage, _ := json.Marshal(message)
	RabbitMQ.SecKillRabbitmq.PublishSimple(string(jsonMessage))
	dto.APIResponse(c, err, request_id.String()+"正在排队中")
	cache.Rdb.Set("status/"+request_id.String(), "1", time.Hour)
	return
}

func checkTime(productID int, secKillList []models.SecKill) bool {
	currTime := time.Now()
	var currSecKill models.SecKill
	for _, secKill := range secKillList {
		if int(secKill.ProductID) == productID {
			currSecKill = secKill
		}
	}

	log.Println(currTime)
	log.Println(currSecKill.StartTime)
	log.Println(currSecKill.EndTime)
	log.Println(currTime.After(currSecKill.StartTime) && currTime.Before(currSecKill.EndTime))
	if currTime.After(currSecKill.StartTime) && currTime.Before(currSecKill.EndTime) {
		return true
	}

	return false
}

// 秒杀后进行支付
func PayForProduct(c *gin.Context) {
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
		if int(product.ID) == id {
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
		dto.APIResponse(c, nil, "已抢到，等待支付")
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
