package router

import (
	"github.com/gin-gonic/gin"
	"server/controller"
	"server/controller/actctrl"
	"server/controller/adminctrl"
	"server/controller/empctrl"
	"server/controller/fruitctrl"
	"server/controller/lgtuserctrl"
	"server/controller/loginctrl"
	"server/controller/regctrl"
	"server/controller/userctrl"
)

// 注册路由
func register(r *gin.Engine) {
	r.POST("/register", regctrl.Register)             // 注册
	r.POST("/login", loginctrl.Login)                 // 登录
	r.PUT("/forget-password", actctrl.ForgetPassword) // 忘记密码

	// 账户管理
	// 内部处理不同权限的功能调用
	accountGroup := r.Group("/account", controller.ValidAuthorization)
	{
		accountGroup.GET("/get", actctrl.GetAccount)                 // 获取个人信息
		accountGroup.PUT("/update", actctrl.Update)                  // 更新个人信息
		accountGroup.GET("/exit", actctrl.Exit)                      // 退出登录
		accountGroup.PUT("/change-password", actctrl.ChangePassword) // 修改账户密码
	}

	// 用户管理
	userGroup := r.Group("/user", controller.ValidAuthorization)
	{
		userGroup.GET("/list", userctrl.List)        // 用户列表
		userGroup.PUT("/update", userctrl.Update)    // 更新用户信息
		userGroup.DELETE("/delete", userctrl.Delete) // 删除用户信息
	}

	// 注销用户管理
	logoutUserGroup := r.Group("/user/logout", controller.ValidAuthorization)
	{
		logoutUserGroup.GET("/list", lgtuserctrl.LogoutList)      // 注销用户列表
		logoutUserGroup.POST("/recover", lgtuserctrl.RecoverUser) // 恢复用户信息
		logoutUserGroup.DELETE("/remove", lgtuserctrl.RemoveUser) // 彻底删除用户
	}

	// 员工管理
	employeeGroup := r.Group("/employee", controller.ValidAuthorization)
	{
		employeeGroup.GET("/list", empctrl.List)           // 员工列表
		employeeGroup.PUT("/update", empctrl.Update)       // 更新员工信息
		employeeGroup.PUT("/promotion", empctrl.Promotion) // 晋升管理员
	}

	// 管理员管理
	adminGroup := r.Group("/admin", controller.ValidAuthorization)
	{
		adminGroup.GET("/list", adminctrl.List) // 管理员列表
	}

	// 水果管理
	fruitGroup := r.Group("/fruit", controller.ValidAuthorization)
	{
		fruitGroup.GET("/list", fruitctrl.List)     // 水果列表
		fruitGroup.GET("/detail", fruitctrl.Detail) // 水果详情
		fruitGroup.POST("/add", fruitctrl.Insert)   // 添加水果
		fruitGroup.DELETE("/del", fruitctrl.Delete) // 删除水果
		fruitGroup.PUT("/update", fruitctrl.Update) // 更新水果信息
	}
}
