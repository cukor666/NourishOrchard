package response

import (
	"github.com/gin-gonic/gin"
	"server/common"
)

// Response 响应对象
type Response struct {
	Code int    `json:"code"` // 状态码
	Msg  string `json:"msg"`  // 消息
	Data any    `json:"data"` // 具体数据
}

// Success 请求成功返回
func Success(c *gin.Context, data any, msg string) {
	c.JSON(200, Response{200, msg, data})
}

// Failed 请求失败返回
func Failed(c *gin.Context, msg string) {
	c.JSON(200, Response{400, msg, 0})
}

// FailedWithCode 请求失败返回，携带有状态码
func FailedWithCode(c *gin.Context, code int, msg string) {
	c.JSON(200, Response{code, msg, 0})
}

func FailedWithError(c *gin.Context, err *common.MyError) {
	c.JSON(err.Code, Response{err.Code, err.Msg, 0})
}
