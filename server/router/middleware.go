package router

import (
	"github.com/gin-gonic/gin"
)

// 中间件
func middleWare(r *gin.Engine) {
	r.Use(corsHandle())
	r.Use(RegisterValid)
}
