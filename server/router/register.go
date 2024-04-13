package router

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/router/routes"
	"sync"
)

// 注册路由
func register(r *gin.Engine) {
	api := r.Group("/api")
	wg := new(sync.WaitGroup)
	wg.Add(10)

	go routes.BaseRoutes(api, wg)       // 基础路由
	go routes.ActRoutes(api, wg)        // 账号管理
	go routes.UserRoutes(api, wg)       // 用户管理
	go routes.LogoutUserRoutes(api, wg) // 注销用户管理
	go routes.EmpRoutes(api, wg)        // 员工管理
	go routes.AdminRoutes(api, wg)      // 管理员管理
	go routes.FruitRoutes(api, wg)      // 水果管理
	go routes.SupplierRoutes(api, wg)   // 供应商管理
	go routes.WarehouseRoutes(api, wg)  // 仓库管理
	go routes.Inventory(api, wg)        // 库存管理

	wg.Wait()
	levellog.Router("全部接口注册完毕")
}
