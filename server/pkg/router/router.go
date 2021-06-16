package router

import (
	"server/pkg/controller"
	"server/pkg/controller/api"

	"github.com/gin-gonic/gin"
)

//Server gin server
type Server struct {
	Engine *gin.Engine
}

//InitRouter ...
func InitRouter() *Server {
	server := new(Server)
	server.Engine = gin.Default()

	registerBaseAPI(server)
	// registerGroupAPI(server)
	registerProductAPI(server)
	return server
}

// registerBaseAPI ...
func registerBaseAPI(server *Server) {
	//Health API
	server.Engine.GET("/health", controller.Health)
	server.Engine.GET("/version", controller.Version)
	//CAS API
	engineGroup := server.Engine.Group("/")
	emptyFunc := func(context *gin.Context) {}
	engineGroup.GET("/login", emptyFunc)
	engineGroup.GET("/logout", emptyFunc)
	engineGroup.GET("/OAuth/GetSession", emptyFunc)
}

// func registerGroupAPI(server *Server) {
// 	engineGroup := server.Engine.Group("/")
// 	engineGroup.GET("/index", api.MultiConsole)
// 	engineGroup.GET("/awsUrl", api.AwsLoginURL)
// 	engineGroup.GET("/aliyunUrl", api.AliyunLoginURL)
// }

func registerProductAPI(server *Server) {
	engineGroup := server.Engine.Group("/")
	//add CAS Auth middleware
	//engineGroup.Use(casClient.Authorize())
	engineGroup.Use(gin.BasicAuth(gin.Accounts{
		"server": "love",
	}))
	engineGroup.GET("/product", api.Products)
	engineGroup.GET("/seckills", api.Seckills)
	engineGroup.GET("/orders", api.Orders)
	engineGroup.POST("/seckill/:productId", api.SecKillProduct)
	engineGroup.POST("/pay/:productId", api.SecKillProduct)
	engineGroup.GET("/status/:requestId", api.GetStatus)
	engineGroup.GET("/status/all", api.GetAllStatus)
}
