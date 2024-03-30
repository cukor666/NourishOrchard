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
	"server/router/valid"
	"server/service/adminsvc"
	"server/service/empsvc"
	"server/service/usersvc"
	"server/utils/promisetool"
)

// ForgetPassword 用户忘记密码
func ForgetPassword(context *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			levellog.Controller(fmt.Sprintf("recover终止，err: %v", err))
			context.Abort()
		}
	}()
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
		if ok := valid.PhoneValidString(req.Phone); !ok {
			levellog.Controller("phone校验失败")
			response.Failed(context, "参数错误，请检查phone的格式")
			return
		}
		err = usersvc.ForgetPassword(req)
	case mc.EMPLOYEE:
		if ok := valid.PhoneValidString(req.Phone); !ok {
			levellog.Controller("phone校验失败")
			response.Failed(context, "参数错误，请检查phone的格式")
			return
		}
		err = empsvc.ForgetPassword(req)
	case mc.ADMIN:
		if ok := valid.EmailValidString(req.Email); !ok {
			levellog.Controller("email校验失败")
			response.Failed(context, "参数错误，请检查email的格式")
			return
		}
		err = adminsvc.ForgetPassword(req)
	default:
		levellog.Controller("未开放")
		err = errors.New("权限未开放")
	}
	if err != nil {
		w := fmt.Sprintf("参数错误，忘记密码接口拒绝请求, err = %s", err.Error())
		levellog.Controller(w)
		response.Failed(context, w)
		return
	}
	response.Success(context, 0, "忘记密码更改密码成功")
}
