package routes

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	"server/controller/adminctrl"
	"sync"
)

// AdminRoutes 管理员管理
func AdminRoutes(rg *gin.RouterGroup, wg *sync.WaitGroup) {
	defer func() {
		levellog.Router("管理员管理接口注册完毕")
		wg.Done()
	}()
	adminGroup := rg.Group("/admin", controller.ValidAuthorization)
	{
		adminGroup.GET("/list", adminctrl.List) // 管理员列表
	}
}
