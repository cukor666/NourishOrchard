package userctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/config"
	"server/controller"
	cm "server/controller/args/claims"
	"server/controller/args/header"
	"server/models"
	mc "server/models/code"
	"server/response"
	"server/service/usersvc"
	"server/utils/promisetool"
)

// Update 更新用户信息
/**
header:
	Bearer Token
	username: CZKJ991706348185
body:
	{
		"id": 7,
		"username": "CZKJ201709005881",
		"name": "吴宣仪大美女",
		"phone": "18577659826",
		"gender": "男",
		"address": "海南省海口市",
		"birthday": "1995-03-16T00:00:00Z"
	}
*/
func Update(context *gin.Context) {
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
	var (
		user models.User
	)
	err = context.ShouldBindJSON(&user)
	if err != nil {
		levellog.Controller("前端数据绑定失败")
		response.Failed(context, "数据绑定失败")
		return
	}
	levellog.Controller(fmt.Sprintf("user: %v", user))
	err = usersvc.Update(user)
	if err != nil {
		levellog.Controller("更新失败")
		response.Failed(context, "系统错误更新失败")
		return
	}
	response.Success(context, user, "更新用户信息成功")
}
