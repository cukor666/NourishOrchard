package router

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	"server/controller/actctrl"
	"server/controller/adminctrl"
	"server/controller/empctrl"
	"server/controller/fruitctrl"
	"server/controller/lgtuserctrl"
	"server/controller/loginctrl"
	"server/controller/regctrl"
	"server/controller/userctrl"
	"sync"
)

// 注册路由
func register(r *gin.Engine) {
	api := r.Group("/api")
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer func() {
			levellog.Router("基接口注册完毕")
			wg.Done()
		}()
		api.POST("/register", regctrl.Register)             // 注册
		api.POST("/login", loginctrl.Login)                 // 登录
		api.PUT("/forget-password", actctrl.ForgetPassword) // 忘记密码
	}()

	wg.Add(1)
	go func() {
		defer func() {
			levellog.Router("账户管理接口注册完毕")
			wg.Done()
		}()
		// 账户管理
		// 内部处理不同权限的功能调用
		accountGroup := api.Group("/account", controller.ValidAuthorization)
		{
			accountGroup.GET("/get", actctrl.GetAccount)                 // 获取个人信息
			accountGroup.PUT("/update", actctrl.Update)                  // 更新个人信息
			accountGroup.GET("/exit", actctrl.Exit)                      // 退出登录
			accountGroup.PUT("/change-password", actctrl.ChangePassword) // 修改账户密码
		}
	}()

	wg.Add(1)
	go func() {
		defer func() {
			levellog.Router("用户管理接口注册完毕")
			wg.Done()
		}()
		// 用户管理
		userGroup := api.Group("/user", controller.ValidAuthorization)
		{
			userGroup.GET("/list", userctrl.List)        // 用户列表
			userGroup.PUT("/update", userctrl.Update)    // 更新用户信息
			userGroup.DELETE("/delete", userctrl.Delete) // 删除用户信息
		}
	}()

	wg.Add(1)
	go func() {
		defer func() {
			levellog.Router("注销用户接口注册完毕")
			wg.Done()
		}()
		// 注销用户管理
		logoutUserGroup := api.Group("/user/logout", controller.ValidAuthorization)
		{
			logoutUserGroup.GET("/list", lgtuserctrl.LogoutList)      // 注销用户列表
			logoutUserGroup.POST("/recover", lgtuserctrl.RecoverUser) // 恢复用户信息
			logoutUserGroup.DELETE("/remove", lgtuserctrl.RemoveUser) // 彻底删除用户
		}
	}()

	wg.Add(1)
	go func() {
		defer func() {
			levellog.Router("员工管理接口注册完毕")
			wg.Done()
		}()
		// 员工管理
		employeeGroup := api.Group("/employee", controller.ValidAuthorization)
		{
			employeeGroup.GET("/list", empctrl.List)           // 员工列表
			employeeGroup.PUT("/update", empctrl.Update)       // 更新员工信息
			employeeGroup.PUT("/promotion", empctrl.Promotion) // 晋升管理员
		}
	}()

	wg.Add(1)
	go func() {
		defer func() {
			levellog.Router("管理员管理接口注册完毕")
			wg.Done()
		}()
		// 管理员管理
		adminGroup := api.Group("/admin", controller.ValidAuthorization)
		{
			adminGroup.GET("/list", adminctrl.List) // 管理员列表
		}
	}()

	wg.Add(1)
	go func() {
		defer func() {
			levellog.Router("水果管理接口注册完毕")
			wg.Done()
		}()
		// 水果管理
		fruitGroup := api.Group("/fruit", controller.ValidAuthorization)
		{
			fruitGroup.GET("/list", fruitctrl.List)     // 水果列表
			fruitGroup.GET("/detail", fruitctrl.Detail) // 水果详情
			fruitGroup.POST("/add", fruitctrl.Insert)   // 添加水果
			fruitGroup.DELETE("/del", fruitctrl.Delete) // 删除水果
			fruitGroup.PUT("/update", fruitctrl.Update) // 更新水果信息
		}
	}()

	wg.Wait()
	levellog.Router("全部接口注册完毕")
}
