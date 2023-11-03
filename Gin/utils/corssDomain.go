package utils

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 解决跨域问题
func CorssDomain() (mwCORS gin.HandlerFunc) {
	mwCORS = cors.New(cors.Config{
		// 运行跨域请求网络，多个网络使用逗号分开，限制使用*
		AllowOrigins: []string{"*"},
		// 允许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "GET", "DELETE"},
		// 允许使用的请求头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		// 显示的请求头
		ExposeHeaders: []string{"Content-Type"},
		// 凭证共享，确定共享
		AllowCredentials: true,
		// 容许跨域的原点网站，可以直接return true就万事大吉了
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		// 超时时间设定
		MaxAge: 24 * time.Hour,
	})
	return
}
