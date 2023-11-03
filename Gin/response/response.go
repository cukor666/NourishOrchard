package response

import "github.com/gin-gonic/gin"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// Success 请求成功返回
func Success(msg string, data any, c *gin.Context) {
	c.JSON(200, Response{200, msg, data})
}

// Failed 请求失败返回
func Failed(msg string, c *gin.Context) {
	c.JSON(200, Response{400, msg, 0})
}
