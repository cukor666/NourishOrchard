package routes

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	"server/controller/empctrl"
	"sync"
)

// EmpRoutes 员工管理
func EmpRoutes(rg *gin.RouterGroup, wg *sync.WaitGroup) {
	defer func() {
		levellog.Router("员工管理接口注册完毕")
		wg.Done()
	}()
	eg := rg.Group("/employee", controller.ValidAuthorization)
	{
		eg.GET("/list", empctrl.List)           // 员工列表
		eg.PUT("/update", empctrl.Update)       // 更新员工信息
		eg.PUT("/promotion", empctrl.Promotion) // 晋升管理员
	}
}
