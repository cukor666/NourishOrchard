package routes

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	"server/controller/fruitctrl"
	"sync"
)

// FruitRoutes 水果管理
func FruitRoutes(rg *gin.RouterGroup, wg *sync.WaitGroup) {
	defer func() {
		levellog.Router("水果管理接口注册完毕")
		wg.Done()
	}()
	fg := rg.Group("/fruit", controller.ValidAuthorization)
	{
		fg.GET("/list", fruitctrl.List)            // 水果列表
		fg.GET("/detail", fruitctrl.Detail)        // 水果详情
		fg.POST("/add", fruitctrl.Insert)          // 添加水果
		fg.DELETE("/del", fruitctrl.Delete)        // 删除水果
		fg.PUT("/update", fruitctrl.Update)        // 更新水果信息
		fg.GET("/origin-pie", fruitctrl.OriginPie) // 产地饼图
	}
}
