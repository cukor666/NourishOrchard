package routes

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	"server/controller/lgtuserctrl"
	"sync"
)

// LogoutUserRoutes 注销用户管理
func LogoutUserRoutes(rg *gin.RouterGroup, wg *sync.WaitGroup) {
	defer func() {
		levellog.Router("注销用户接口注册完毕")
		wg.Done()
	}()
	lug := rg.Group("/user/logout", controller.ValidAuthorization)
	{
		lug.GET("/list", lgtuserctrl.LogoutList)      // 注销用户列表
		lug.POST("/recover", lgtuserctrl.RecoverUser) // 恢复用户信息
		lug.DELETE("/remove", lgtuserctrl.RemoveUser) // 彻底删除用户
	}
}
