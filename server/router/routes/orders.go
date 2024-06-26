package routes

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	"server/controller/orderctrl"
	"sync"
)

func OrdersRoutes(rg *gin.RouterGroup, wg *sync.WaitGroup) {
	defer func() {
		levellog.Router("订单管理接口注册完毕")
		wg.Done()
	}()
	group := rg.Group("/orders", controller.ValidAuthorization)
	{
		group.GET("/list", orderctrl.List)     // 获取订单列表
		group.POST("/add", orderctrl.Insert)   // 新增订单
		group.DELETE("/del", orderctrl.Delete) // 删除订单
		group.PUT("/update", orderctrl.Update) // 更新订单
	}
	userGroup := rg.Group("/user-orders") // 用户端订单接口
	{
		userGroup.GET("/list", orderctrl.List) // 获取用户订单列表
	}
}
