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
	wg.Add(8)

	go routes.BaseRoutes(api, wg)
	go routes.ActRoutes(api, wg)
	go routes.UserRoutes(api, wg)
	go routes.LogoutUserRoutes(api, wg)
	go routes.EmpRoutes(api, wg)
	go routes.AdminRoutes(api, wg)
	go routes.FruitRoutes(api, wg)
	go routes.SupplierRoutes(api, wg)

	wg.Wait()
	levellog.Router("全部接口注册完毕")
}
