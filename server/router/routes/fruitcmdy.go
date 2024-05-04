package routes

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller/fruitcmdyctrl"
	"sync"
)

// FruitCmdyRoutes 水果商品管理
func FruitCmdyRoutes(rg *gin.RouterGroup, wg *sync.WaitGroup) {
	defer func() {
		levellog.Router("水果商品管理路由初始化完成")
		wg.Done()
	}()
	group := rg.Group("/fruitcmdy")
	{
		group.GET("/list", fruitcmdyctrl.List) // 获取水果商品列表
	}
}
