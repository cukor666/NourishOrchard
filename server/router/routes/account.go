package routes

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	"server/controller/actctrl"
	"sync"
)

// ActRoutes 账户管理
// 内部处理不同权限的功能调用
func ActRoutes(rg *gin.RouterGroup, wg *sync.WaitGroup) {
	defer func() {
		levellog.Router("账户管理接口注册完毕")
		wg.Done()
	}()
	ag := rg.Group("/account", controller.ValidAuthorization)
	{
		ag.GET("/get", actctrl.GetAccount)                 // 获取个人信息
		ag.PUT("/update", actctrl.Update)                  // 更新个人信息
		ag.GET("/exit", actctrl.Exit)                      // 退出登录
		ag.PUT("/change-password", actctrl.ChangePassword) // 修改账户密码
		ag.GET("/cnt", actctrl.Count)                      // 获取账户数量
	}
}
