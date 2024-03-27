package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/request"
	"server/response"
	"server/service"
)

// 数据校验
func (rc *RegisterController) validation(rg request.RegisterRequest) bool {
	return len(rg.Name) >= 2 && len(rg.Name) <= 20 && len(rg.Password) >= 3 && len(rg.Password) <= 20
}

// Register 注册
/**
header: 空
body: application/json
{
    "name": "李丽梅",
    "password": "lilimei",
    "promise": 2458,
    "gender": "女",
    "phone": "13746510502",
    "address": "广东省深圳市",
    "birthday": "1996-08-26T15:03:10Z"
}
*/
func (rc *RegisterController) Register(context *gin.Context) {
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
	ok := rc.validation(registerRequest)
	if !ok {
		response.Failed(context, "参数错误")
		return
	}
	// 数据校验通过，交给业务层
	registerResponse, ok = service.RegisterService{}.Register(registerRequest)
	if !ok {
		levelLog("注册失败")
		response.FailedWithCode(context, 500, "注册失败")
		return
	}
	response.Success(context, registerResponse, "绑定成功")
}
