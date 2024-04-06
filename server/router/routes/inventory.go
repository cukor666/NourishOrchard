package routes

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	"server/controller/ivtctrl"
	"sync"
)

func Inventory(rg *gin.RouterGroup, wg *sync.WaitGroup) {
	defer func() {
		levellog.Router("库存管理接口注册完毕")
		wg.Done()
	}()
	group := rg.Group("/ivt", controller.ValidAuthorization)
	{
		group.GET("/list", ivtctrl.List)     // 库存列表
		group.POST("/add", ivtctrl.Insert)   // 添加库存信息
		group.DELETE("/del", ivtctrl.Delete) // 删除库存信息
		group.PUT("/update", ivtctrl.Update) // 更新库存信息
	}
}
