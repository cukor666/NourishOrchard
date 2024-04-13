package actctrl

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller/spliterr"
	mc "server/models/code"
	"server/response"
	"server/service/adminsvc"
	"server/service/empsvc"
	"server/service/usersvc"
	"server/utils/promisetool"
)

// ForgetPassword 用户忘记密码
func ForgetPassword(context *gin.Context) {
	var (
		req    forgetPwdReq
		err    error
		myCode = "1024" // 校验短信验证码
	)

	if err = context.ShouldBindJSON(&req); err != nil {
		str := spliterr.GetErrMsg(err.Error())
		levellog.Controller(fmt.Sprintf("数据绑定失败， err: %s", str))
		response.Failed(context, str)
		return
	}

	if req.Code != myCode {
		levellog.Controller(fmt.Sprintf("验证码错误应该得到：%s，却得到%s", myCode, req.Code))
		response.Failed(context, "验证码错误")
		return
	}

	promise := promisetool.ToInt(req.Promise)
	switch promise {
	case mc.USER:
		err = usersvc.ForgetPassword(req.toUser(), req.Password)
	case mc.EMPLOYEE:
		err = empsvc.ForgetPassword(req.toEmp(), req.Password)
	case mc.ADMIN:
		err = adminsvc.ForgetPassword(req.toAdmin(), req.Password)
	default:
		levellog.Controller("权限错误，忘记密码接口拒绝请求")
		err = errors.New("权限错误，忘记密码接口拒绝请求")
	}
	if err != nil {
		w := fmt.Sprintf("参数错误，忘记密码接口拒绝请求, err = %s", err.Error())
		levellog.Controller(w)
		response.Failed(context, w)
		return
	}
	response.Success(context, 0, "忘记密码更改密码成功")
}
