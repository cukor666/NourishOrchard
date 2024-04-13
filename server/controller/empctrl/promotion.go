package empctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	"server/response"
	"server/service/actsvc"
	"server/service/empsvc"
	"server/utils/pwdtool"
)

// Promotion 晋升管理员
func Promotion(context *gin.Context) {
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限校验失败")
		response.Failed(context, "权限校验失败")
		return
	}
	reqUsername := controller.GetUsername(claims)
	reqPromise := controller.GetPromise(claims)

	var req promoteReq
	if err = context.ShouldBindJSON(&req); err != nil {
		w := fmt.Sprintf("参数格式错误: %s", err.Error())
		levellog.Controller(fmt.Sprintf("err: %s", w))
		response.Failed(context, w)
		return
	}

	// 校验密码
	account, err := actsvc.Get(reqUsername, reqPromise)
	if err != nil {
		levellog.Controller("获取账号信息失败")
		response.Failed(context, "获取账号信息失败")
		return
	}

	if !pwdtool.PwdOK(account.Password, req.Password) {
		levellog.Controller(fmt.Sprintf("请求账号%s密码错误，拒绝请求", reqUsername))
		response.Failed(context, "密码错误")
		return
	}

	if err = empsvc.Promotion(req.Username, req.Name, req.Email, req.Mark); err != nil {
		levellog.Controller("晋升管理员失败")
		response.Failed(context, "晋升管理员失败")
		return
	}
	response.Success(context, 0, "晋升管理员成功")
}
