package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/router/routes"
	"sync"
)

// 注册路由
func register(r *gin.Engine) {
	// 注册接口组
	api := r.Group("/api")

	// 注册接口
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 退出时取消注册

	var wg sync.WaitGroup

	// 注册接口
	routeList := []func(*gin.RouterGroup, *sync.WaitGroup){
		routes.BaseRoutes,       // 基础路由
		routes.ActRoutes,        // 账户管理
		routes.UserRoutes,       // 用户管理
		routes.LogoutUserRoutes, // 注销用户管理
		routes.EmpRoutes,        // 员工管理
		routes.AdminRoutes,      // 管理员管理
		routes.FruitRoutes,      // 水果管理
		routes.SupplierRoutes,   // 供应商管理
		routes.WarehouseRoutes,  // 仓库管理
		routes.Inventory,        // 库存管理
		routes.OrdersRoutes,     // 订单管理
	}

	wg.Add(len(routeList))

	// 启动goroutine
	for _, route := range routeList {
		go route(api, &wg)
	}

	// 等待所有goroutine完成
	go func() {
		wg.Wait()
		cancel()
	}()

	// 等待退出信号
	<-ctx.Done()

	levellog.Router("全部接口注册完毕")
}
