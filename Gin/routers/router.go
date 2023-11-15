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

	r.GET("/", utils.JWTAuthMiddleware(), api.AuthHandler) // 鉴权
	r.GET("/captcha", api.Captcha)                         // 验证码
	r.POST("/user-login", api.AuthHandler)                 // 登录时使用，新版，带鉴权
	r.POST("/captcha/verify", api.Verify)                  // 验证码校验
	r.GET("/promise/:name", api.GetUserPromise)            // 获取用户权限

	// 路由组
	// 普通用户
	userGroup := r.Group("/user")
	{
		// GET
		userGroup.GET("/list", api.UserList)           // 展示用户列表
		userGroup.GET("/:id", api.FindUser)            // 根据ID查找用户
		userGroup.GET("/", api.SeeUserInfo)            // 根据Name查找用户
		userGroup.GET("/struct", api.FindUserByStruct) // 根据User结构体查询用户
		userGroup.GET("/count", api.GetUserCounnt)     // 获取用户个数

		// POST
		userGroup.POST("/add", api.UserAdd) // 添加用户，前端注册的时候使用到这个

		// PUT
		userGroup.PUT("/update", api.UpdateUserInfo) // 更新用户信息，在用户列表的编辑按钮使用

		// DELETE
		userGroup.DELETE("/delete", api.DeleteUser) // 删除用户信息，在用户列表的删除按钮使用
	}

	// 管理员
	adminGroup := r.Group("/admin")
	{
		// GET
		adminGroup.GET("/list", api.AdminList) // 分页查询所有管理员

		// POST

		// PUT
		adminGroup.PUT("/adds", api.AddAdmin)      // 虽然是增加操作，但实际上只是将权限更新为管理员身份
		adminGroup.PUT("/delete", api.DeleteAdmin) // 虽然是删除操作，但实际上只是将权限更新为普通用户权限
		// DELETE

	}
}
