package lgtuserctrl

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
	"server/service/lgtusersvc"
	"server/utils/promisetool"
)

// RecoverUser 恢复用户信息
/**
header:
	Bearer Token
	username: CZKJ991706348185
body:
	{
		"username": "CZKJ991706340782"
	}
*/
func RecoverUser(context *gin.Context) {
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
	// 获取请求方的权限，如果是普通用户权限则无权访问
	promise := promisetool.ToInt(claims[cm.Promise].(string))
	if promise < mc.ADMIN {
		levellog.Controller("权限不够，更新用户信息")
		response.Failed(context, "权限不够，无权操作")
		return
	}

	type tempReq struct {
		Username string `json:"username" binding:"required,username"`
	}
	var req tempReq
	err = context.ShouldBindJSON(&req)
	if err != nil {
		str := spliterr.GetErrMsg(fmt.Sprintf("参数绑定失败，err: %s", err.Error()))
		levellog.Controller(str)
		response.Failed(context, str)
		return
	}

	user, err := lgtusersvc.RecoverUser(req.Username)
	if err != nil {
		levellog.Controller(fmt.Sprintf("恢复用户信息失败，user = %v", user))
		response.Failed(context, "恢复用户信息失败")
		return
	}
	response.Success(context, user, "恢复用户信息成功")
}
