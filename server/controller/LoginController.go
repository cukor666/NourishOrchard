package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common"
	"server/common/levellog"
	"server/request"
	"server/response"
	"server/service"
	"server/utils/promisetool"
)

// 登录参数校验
func (lc *LoginController) validation(req request.LoginRequest) bool {
	if req.Username == "" || req.Password == "" || promisetool.ToInt(req.Promise) == 0 {
		return false
	}
	return true
}

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
func (lc *LoginController) Login(context *gin.Context) {
	var (
		loginRequest request.LoginRequest
		myErr        *common.MyError
	)
	if context.ShouldBind(&loginRequest) != nil {
		levellog.Controller("绑定失败")
		response.Failed(context, "参数错误")
		return
	}
	ok := lc.validation(loginRequest)
	if !ok {
		levellog.Controller("参数校验不通过")
		response.Failed(context, "参数校验不通过")
		return
	}
	// 登录业务
	token, err := service.LoginService{}.Login(loginRequest)
	if errors.As(err, &myErr) {
		levellog.Controller("登录失败")
		response.FailedWithError(context, myErr)
		return
	}
	levellog.Controller(fmt.Sprintf("token: %v", token))
	response.Success(context, token, "登录成功")
}
