package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/config"
	"server/utils"
)

// Start 启动gin
func Start() {
	r := gin.Default()
	//middleWare(r) // 中间件
	register(r) // 注册路由
	if err := r.Run(fmt.Sprintf(":%d", config.GetConfig().SystemConfig.Port)); err != nil {
		panic(fmt.Sprintf("启动失败,错误信息: myerr = %v", err))
	}
}

// 中间件
//func middleWare(r *gin.Engine) {
//	//r.Use(utils.CorssDomain())                                   // 解决跨域问题
//	//r.Use(utils.Session(config.GetConfig().SystemConfig.Secret)) // 配置session
//}

// 注册路由
func register(r *gin.Engine) {
	r.POST("/register", registerController.Register)
	r.POST("/login", loginController.Login)

	// 账户管理
	accountGroup := r.Group("/account", utils.ValidAuthorization)
	{
		accountGroup.GET("/get", accountController.GetAccount)
	}
}

