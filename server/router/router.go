package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/config"
	"time"
)

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
