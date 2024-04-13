package regctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/response"
	"server/service/regsvc"
)

// Register 注册
func Register(context *gin.Context) {
	var (
		req registerRequest
		res registerResponse
	)
	if err := context.ShouldBindJSON(&req); err != nil {
		levellog.Controller(fmt.Sprintf("绑定参数失败, err: %s", err.Error()))
		response.Failed(context, "绑定参数失败")
		return
	}

	// 调用注册服务
	result, ok := regsvc.Register(req.Deconstruct())
	if !ok {
		levellog.Controller("注册失败")
		response.FailedWithCode(context, 500, "注册失败")
		return
	}
	res.set(result)
	response.Success(context, res, "绑定成功")
}
