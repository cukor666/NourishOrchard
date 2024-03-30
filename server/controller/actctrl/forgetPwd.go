package actctrl

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller/spliterr"
	mc "server/models/code"
	"server/request"
	"server/response"
	"server/service/adminsvc"
	"server/service/empsvc"
	"server/service/usersvc"
	"server/utils/promisetool"
)

// ForgetPassword 用户忘记密码
func ForgetPassword(context *gin.Context) {
	var (
		req request.ForgetPwdReq
		err error
	)

	err = context.ShouldBindJSON(&req)
	if err != nil {
		str := spliterr.GetErrMsg(err.Error())
		levellog.Controller(fmt.Sprintf("数据绑定失败， err: %s", str))
		response.Failed(context, str)
		return
	}

	// 校验短信验证码
	myCode := "1024"
	if req.Code != myCode {
		levellog.Controller(fmt.Sprintf("验证码错误应该得到：%s，却得到%s", myCode, req.Code))
		response.Failed(context, "验证码错误")
		return
	}

	promise := promisetool.ToInt(req.Promise)
	switch promise {
	case mc.USER:
		err = usersvc.ForgetPassword(req)
	case mc.EMPLOYEE:
		err = empsvc.ForgetPassword(req)
	case mc.ADMIN:
		err = adminsvc.ForgetPassword(req)
	default:
		levellog.Controller("未开放")
		err = errors.New("权限未开放")
	}
	if err != nil {
		levellog.Controller("服务器错误，忘记密码接口拒绝请求")
		response.Failed(context, "服务器错误，忘记密码接口拒绝请求")
		return
	}
	response.Success(context, 0, "忘记密码更改密码成功")
}
