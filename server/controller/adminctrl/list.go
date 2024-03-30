package adminctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/common/simpletool"
	"server/config"
	"server/controller"
	cm "server/controller/args/claims"
	"server/controller/args/header"
	"server/controller/spliterr"
	mc "server/models/code"
	"server/response"
	"server/service/adminsvc"
	"server/utils/promisetool"
)

// List 查询管理员列表
/**
header:
	Bearer Token
	username: CZKJ991706348185
params:
	query:
		pageSize: 3
		pageNum: 2
*/
func List(context *gin.Context) {
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
		levellog.Controller("权限不够，无权访问用户列表")
		response.Failed(context, "权限不够，无权访问")
		return
	}
	var p simpletool.Page

	err = context.ShouldBindQuery(&p)
	if err != nil {
		msg := spliterr.GetErrMsg(err.Error())
		w := fmt.Sprintf("参数校验失败, err: %s", msg)
		levellog.Controller(w)
		response.Failed(context, w)
		return
	}

	admins, total, err := adminsvc.ListWithPage(p)
	if err != nil {
		levellog.Controller("服务器端错误，查询失败")
		response.Failed(context, "服务器端错误")
		return
	}
	response.Success(context, gin.H{
		"admins": admins,
		"total":  total,
	}, "查询管理员列表成功")
}
