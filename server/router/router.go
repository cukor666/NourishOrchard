package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"server/config"
	"server/controller"
	"time"
)

func newLogFile() {
	// 创建一个新的日志文件
	file, err := os.OpenFile("log/app.log."+time.Now().Format("2006-01-02"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// 将文件设置为 log 包的输出目标
	log.SetOutput(file)
}

// Start 启动gin
func Start() {
	ticker := time.NewTicker(24 * time.Hour)
	newLogFile()

	r := gin.Default()
	middleWare(r) // 中间件
	register(r)   // 注册路由

	// 启动定时器
	go func() {
		for range ticker.C {
			newLogFile()
		}
	}()

	if err := r.Run(fmt.Sprintf(":%d", config.GetConfig().SystemConfig.Port)); err != nil {
		panic(fmt.Sprintf("启动失败,错误信息: myerr = %v", err))
	}
}

// 中间件
func middleWare(r *gin.Engine) {
	// 配置CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     domains[:], // 允许的域名列表，使用*允许所有域名
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "username"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

// 注册路由
func register(r *gin.Engine) {
	r.POST("/register", controller.RegisterController{}.Register)            // 注册
	r.POST("/login", controller.LoginController{}.Login)                     // 登录
	r.PUT("/forget-password", controller.AccountController{}.ForgetPassword) // 忘记密码

	// 账户管理
	// 内部处理不同权限的功能调用
	accountGroup := r.Group("/account", controller.ValidAuthorization)
	{
		accountGroup.GET("/get", controller.AccountController{}.GetAccount)                 // 获取个人信息
		accountGroup.PUT("/update", controller.AccountController{}.Update)                  // 更新个人信息
		accountGroup.GET("/exit", controller.AccountController{}.Exit)                      // 退出登录
		accountGroup.PUT("/change-password", controller.AccountController{}.ChangePassword) // 修改账户密码
	}

	// 用户管理
	userGroup := r.Group("/user", controller.ValidAuthorization)
	{
		userGroup.GET("/list", controller.UserController{}.List)              // 用户列表
		userGroup.PUT("/update", controller.UserController{}.Update)          // 更新用户信息
		userGroup.DELETE("/delete", controller.UserController{}.Delete)       // 删除用户信息
		userGroup.GET("/logout-list", controller.UserController{}.LogoutList) // 注销用户列表
		userGroup.POST("/recover", controller.UserController{}.RecoverUser)   // 恢复用户信息
		userGroup.DELETE("/remove", controller.UserController{}.RemoveUser)   // 彻底删除用户
	}

	// 员工管理
	employeeGroup := r.Group("/employee", controller.ValidAuthorization)
	{
		employeeGroup.GET("/list", controller.EmployeeController{}.List)     // 员工列表
		employeeGroup.PUT("/update", controller.EmployeeController{}.Update) // 更新员工信息
	}

	// 管理员管理
	adminGroup := r.Group("/admin", controller.ValidAuthorization)
	{
		adminGroup.GET("/list", controller.AdminController{}.List) // 管理员列表
	}

	// 水果管理
	fruitGroup := r.Group("/fruit", controller.ValidAuthorization)
	{
		fruitGroup.GET("/list", controller.FruitController{}.List) // 水果列表
	}
}
