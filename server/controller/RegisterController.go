package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/request"
	"server/response"
)

// 数据校验
func (r RegisterController) validation(rg request.RegisterRequest) bool {
	return len(rg.Name) >= 2 && len(rg.Name) <= 20 && len(rg.Password) >= 3 && len(rg.Password) <= 20
}

// Register 注册
func (r RegisterController) Register(context *gin.Context) {
	var (
		registerRequest  request.RegisterRequest
		registerResponse response.RegisterResponse
	)
	if context.ShouldBind(&registerRequest) != nil {
		levelLog("绑定参数失败")
		response.Failed(context, "绑定参数失败")
		return
	}
	levelLog("绑定成功")
	levelLog(fmt.Sprintf("%v", registerRequest))
	ok := r.validation(registerRequest)
	if !ok {
		response.Failed(context, "参数错误")
		return
	}
	// 数据校验通过，交给业务层
	registerResponse, ok = registerService.Register(registerRequest)
	if !ok {
		levelLog("注册失败")
		response.FailedWithCode(context, 500, "注册失败")
		return
	}
	response.Success(context, registerResponse, "绑定成功")
}
