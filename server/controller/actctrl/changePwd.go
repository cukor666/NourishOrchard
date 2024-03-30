package actctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/config"
	"server/controller"
	cm "server/controller/args/claims"
	"server/controller/args/header"
	"server/controller/spliterr"
	mc "server/models/code"
	"server/response"
	"server/service/actsvc"
	"server/utils/promisetool"
)

// ChangePassword 修改账户密码
func ChangePassword(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := controller.GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levellog.Controller("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		levellog.Controller("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	username := claims[cm.Username]
	promise := promisetool.ToInt(claims[cm.Promise].(string))
	type pwdType struct {
		OldPassword string `json:"oldPassword" binding:"required,min=3,max=20"`
		NewPassword string `json:"newPassword" binding:"required,min=3,max=20"`
	}
	var pwd pwdType
	err = context.ShouldBindJSON(&pwd)
	if err != nil {
		str := spliterr.GetErrMsg(fmt.Sprintf("校验前端数据失败, err: %s", err.Error()))
		levellog.Controller(str)
		response.Failed(context, str)
		return
	}

	switch promise {
	case mc.USER, mc.ADMIN, mc.EMPLOYEE:
		err = actsvc.ChangePassword(username.(string), pwd.OldPassword, pwd.NewPassword)
		if err != nil {
			levellog.Controller("服务器错误修改密码失败")
			response.Failed(context, "服务器错误修改密码失败")
			return
		}
		response.Success(context, username, "修改密码成功")
	default:
		levellog.Controller(fmt.Sprintf("权限不正确，promise: %d", promise))
		response.Failed(context, "权限不正确")
		return
	}
}
