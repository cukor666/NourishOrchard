package router

import (
	"github.com/gin-gonic/gin"
	"server/controller"
)

// 注册路由
func register(r *gin.Engine) {
	r.POST("/register", registerController.Register)            // 注册
	r.POST("/login", loginController.Login)                     // 登录
	r.PUT("/forget-password", accountController.ForgetPassword) // 忘记密码

	// 账户管理
	// 内部处理不同权限的功能调用
	accountGroup := r.Group("/account", controller.ValidAuthorization)
	{
		accountGroup.GET("/get", accountController.GetAccount)                 // 获取个人信息
		accountGroup.PUT("/update", accountController.Update)                  // 更新个人信息
		accountGroup.GET("/exit", accountController.Exit)                      // 退出登录
		accountGroup.PUT("/change-password", accountController.ChangePassword) // 修改账户密码
	}

	// 用户管理
	userGroup := r.Group("/user", controller.ValidAuthorization)
	{
		userGroup.GET("/list", userController.List)              // 用户列表
		userGroup.PUT("/update", userController.Update)          // 更新用户信息
		userGroup.DELETE("/delete", userController.Delete)       // 删除用户信息
		userGroup.GET("/logout-list", userController.LogoutList) // 注销用户列表
		userGroup.POST("/recover", userController.RecoverUser)   // 恢复用户信息
		userGroup.DELETE("/remove", userController.RemoveUser)   // 彻底删除用户
	}

	// 员工管理
	employeeGroup := r.Group("/employee", controller.ValidAuthorization)
	{
		employeeGroup.GET("/list", employeeController.List)           // 员工列表
		employeeGroup.PUT("/update", employeeController.Update)       // 更新员工信息
		employeeGroup.PUT("/promotion", employeeController.Promotion) // 晋升管理员
	}

	// 管理员管理
	adminGroup := r.Group("/admin", controller.ValidAuthorization)
	{
		adminGroup.GET("/list", adminController.List) // 管理员列表
	}

	// 水果管理
	fruitGroup := r.Group("/fruit", controller.ValidAuthorization)
	{
		fruitGroup.GET("/list", fruitController.List)     // 水果列表
		fruitGroup.GET("/detail", fruitController.Detail) // 水果详情
		fruitGroup.POST("/add", fruitController.Insert)   // 添加水果
		fruitGroup.DELETE("/del", fruitController.Delete) // 删除水果
		fruitGroup.PUT("/update", fruitController.Update) // 更新水果信息
	}
}
