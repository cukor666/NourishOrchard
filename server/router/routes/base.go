package routes

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller/actctrl"
	"server/controller/loginctrl"
	"server/controller/regctrl"
	"sync"
)

// BaseRoutes 基路由
func BaseRoutes(rg *gin.RouterGroup, wg *sync.WaitGroup) {
	defer func() {
		levellog.Router("基接口注册完毕")
		wg.Done()
	}()
	rg.POST("/register", regctrl.Register)             // 注册
	rg.POST("/login", loginctrl.Login)                 // 登录
	rg.PUT("/forget-password", actctrl.ForgetPassword) // 忘记密码
}
