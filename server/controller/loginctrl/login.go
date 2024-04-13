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
		loginRequest request.LoginRequest // 定义登录请求结构
		myErr        *common.MyError      // 定义自定义错误
	)
	if err := context.ShouldBind(&loginRequest); err != nil { // 校验参数是否绑定成功
		levellog.Controller(spliterr.GetErrMsg(err.Error())) // 记录日志
		response.Failed(context, "参数校验失败")                   // 返回参数校验失败响应
		return
	}

	// 登录业务
	token, err := loginsvc.Login(loginRequest) // 执行登录操作
	if errors.As(err, &myErr) {                // 判断登录错误类型
		levellog.Controller("登录失败")              // 记录日志
		response.FailedWithError(context, myErr) // 返回登录错误响应
		return
	}
	levellog.Controller(fmt.Sprintf("token: %v", token)) // 记录日志
	response.Success(context, token, "登录成功")             // 返回登录成功响应
}
