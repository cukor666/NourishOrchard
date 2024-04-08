package routes

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	"server/controller/pchactrl"
	"sync"
)

func PurchaseRoutes(rg *gin.RouterGroup, wg *sync.WaitGroup) {
	defer func() {
		levellog.Router("采购管理接口注册完毕")
		wg.Done()
	}()
	group := rg.Group("/purchase", controller.ValidAuthorization)
	{
		group.GET("/list", pchactrl.List) // 采购列表
	}
}
