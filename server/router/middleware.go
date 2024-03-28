package router

import (
	"github.com/gin-gonic/gin"
)

// 中间件
func middleWare(r *gin.Engine) {
	// 配置CORS
	r.Use(corsHandle())
	// 注册自定义校验器
	r.Use(RegisterValid())
}
