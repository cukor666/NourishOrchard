package userctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	mc "server/models/code"
	"server/response"
	"server/service/usersvc"
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
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限校验失败")
		response.Failed(context, "权限校验失败")
		return
	}
	// 获取请求方的权限，如果是普通用户权限则无权访问
	promise := controller.GetPromise(claims)
	if promise < mc.ADMIN {
		levellog.Controller("权限不够，更新用户信息")
		response.Failed(context, "权限不够，无权操作")
		return
	}
	var user UpdateUserReq
	err = context.ShouldBindJSON(&user)
	if err != nil {
		w := fmt.Sprintf("前端数据绑定失败, err: %s", err.Error())
		levellog.Controller(w)
		response.Failed(context, w)
		return
	}
	levellog.Controller(fmt.Sprintf("user: %v", user))
	err = usersvc.Update(user.ToUser())
	if err != nil {
		levellog.Controller(fmt.Sprintf("更新失败, err: %s", err.Error()))
		response.Failed(context, "系统错误更新失败")
		return
	}
	response.Success(context, user, "更新用户信息成功")
}
