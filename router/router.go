package router

import (
	"armor_plate/controller"
	_ "armor_plate/docs"
	"armor_plate/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == "dev" {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	v1 := r.Group("/server/v1")
	//健康检测
	v1.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "server is pong")
	})
	v1.POST("/user/login", controller.LoginHandler)
	v1.POST("/user/register", controller.RegisterHandler)
	v1.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		
		v1.POST("/user/setUserAuth", controller.SetUserAuthHandler) //设置用户角色
		v1.POST("/user/createAuth", controller.CreateAuthHandler)   //创建角色权限
		v1.POST("/user/refreshToken", controller.RefreshTokenHandler)
		v1.POST("/user/updateUser", controller.UpdateHandler) //更新用户信息
		v1.GET("/user/getUserList", controller.GetEmpListHandler)
		
		v1.POST("/depot/creat_depot", controller.CreateDepotHandler) //添加仓库信息
		
		v1.POST("/armor/add_product", controller.AddProductHandler)        //添加货物信息
		v1.POST("/armor/setPath", controller.SetProductPathHandler)        //货物入仓
		v1.GET("/armor/getProductPathList", controller.GetPathListHandler) //获取货物存储位置列表
		
		v1.POST("/order/create_order", controller.CreateOrderHandler)    //创建订单信息
		v1.POST("/order/deliver_order", controller.DeliveryOrderHandler) //交付订单业务
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
