package routers

import (
	"Gin/api"
	"Gin/utils"

	"github.com/gin-gonic/gin"
)

// 启动gin
func Start(port string) {
	r := gin.Default()
	middleWare(r) // 中间件
	register(r)   // 注册路由
	r.Run(":" + port)
}

// 中间件
func middleWare(r *gin.Engine) {
	r.Use(utils.CorssDomain())       // 解决跨域问题
	r.Use(utils.Session("cukor.cn")) // 配置session
}

// 注册路由
func register(r *gin.Engine) {
	// GET
	r.GET("/user-list", api.UserList)                          // 展示用户列表
	r.GET("/user/:id", api.FindUser)                           // 根据ID查找用户
	r.GET("/user", api.SeeUserInfo)                            // 根据Name查找用户
	r.GET("/user-struct", api.FindUserByStruct)                // 根据User结构体查询用户
	r.GET("/home", utils.JWTAuthMiddleware(), api.AuthHandler) // 鉴权
	r.GET("/captcha", api.Captcha)                             // 验证码

	// POST
	r.POST("/addUser", api.UserAdd) // 添加用户，前端注册的时候使用到这个
	// r.POST("/user-login", api.UserLogin) // 登录时使用，旧版，不带鉴权
	r.POST("/user-login", api.AuthHandler) // 登录时使用，新版，带鉴权
	// r.POST("/auth", api.AuthHandler)     // jwt鉴权
	r.POST("/captcha/verify", api.Verify) // 验证码校验

	// PUT
	r.PUT("/update-user-info", api.UpdateUserInfo) // 更新用户信息，在用户列表的编辑按钮使用

	// DELETE
	r.DELETE("/delete-user", api.DeleteUser) // 删除用户信息，在用户列表的删除按钮使用
}
