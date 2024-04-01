package routes

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	"server/controller/supctrl"
	"sync"
)

// SupplierRoutes 供应商管理
func SupplierRoutes(rg *gin.RouterGroup, wg *sync.WaitGroup) {
	defer func() {
		levellog.Router("供应商管理接口注册完毕")
		wg.Done()
	}()
	sg := rg.Group("/supplier", controller.ValidAuthorization)
	{
		sg.GET("/list", supctrl.List)        // 供应商列表
		sg.POST("/add", supctrl.Insert)      // 添加供应商信息
		sg.DELETE("/delete", supctrl.Delete) // 删除供应商信息
		sg.PUT("/update", supctrl.Update)    // 更新供应商信息
	}
}
