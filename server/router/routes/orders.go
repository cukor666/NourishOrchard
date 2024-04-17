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
		group.GET("/list", orderctrl.List) // 获取订单列表
	}
}
