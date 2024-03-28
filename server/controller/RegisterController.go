package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller/spliterr"
	"server/request"
	"server/response"
	"server/service"
)

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
	var registerRequest request.RegisterRequest
	if err := context.ShouldBindJSON(&registerRequest); err != nil {
		msg := spliterr.GetErrMsg(err.Error())
		levellog.Controller(fmt.Sprintf("绑定参数失败, err: %s", msg))
		response.Failed(context, "绑定参数失败")
		return
	}

	// 数据校验通过，交给业务层
	registerResponse, ok := service.RegisterService{}.Register(registerRequest)
	if !ok {
		levellog.Controller("注册失败")
		response.FailedWithCode(context, 500, "注册失败")
		return
	}
	response.Success(context, registerResponse, "绑定成功")
}
