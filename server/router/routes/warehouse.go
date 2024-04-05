package routes

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	"server/controller/whctrl"
	"sync"
)

// WarehouseRoutes 仓库管理
func WarehouseRoutes(rg *gin.RouterGroup, wg *sync.WaitGroup) {
	defer func() {
		levellog.Router("仓库管理接口注册完毕")
		wg.Done()
	}()
	g := rg.Group("/warehouse", controller.ValidAuthorization)
	{
		g.GET("/list", whctrl.List)     // 仓库列表
		g.POST("/add", whctrl.Insert)   // 插入新仓库
		g.DELETE("/del", whctrl.Delete) // 删除仓库
		g.PUT("/update", whctrl.Update) // 更新仓库信息
	}
}
