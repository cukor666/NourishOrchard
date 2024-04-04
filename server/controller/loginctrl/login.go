package loginctrl

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common"
	"server/common/levellog"
	"server/controller/spliterr"
	"server/request"
	"server/response"
	"server/service/loginsvc"
)

// Login 登录接口
/**
header: 空
body: application/json
	{
		"username": "CZKJ901704561805",
		"password": "123456",
		"promise": "user"
	}
*/
func Login(context *gin.Context) {
	var (
		loginRequest request.LoginRequest
		myErr        *common.MyError
		err          error
	)
	if err = context.ShouldBind(&loginRequest); err != nil {
		levellog.Controller(spliterr.GetErrMsg(err.Error()))
		response.Failed(context, "参数校验失败")
		return
	}

	// 登录业务
	token, err := loginsvc.Login(loginRequest)
	if errors.As(err, &myErr) {
		levellog.Controller("登录失败")
		response.FailedWithError(context, myErr)
		return
	}
	levellog.Controller(fmt.Sprintf("token: %v", token))
	response.Success(context, token, "登录成功")
}