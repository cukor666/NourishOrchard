package routes

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	"server/controller/userctrl"
	"sync"
)

// UserRoutes 用户管理
func UserRoutes(rg *gin.RouterGroup, wg *sync.WaitGroup) {
	defer func() {
		levellog.Router("用户管理接口注册完毕")
		wg.Done()
	}()
	ug := rg.Group("/user", controller.ValidAuthorization)
	{
		ug.GET("/list", userctrl.List)        // 用户列表
		ug.PUT("/update", userctrl.Update)    // 更新用户信息
		ug.DELETE("/delete", userctrl.Delete) // 删除用户信息
	}
}
