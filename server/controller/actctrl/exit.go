package actctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/config"
	"server/controller"
	cm "server/controller/args/claims"
	"server/controller/args/header"
	mc "server/models/code"
	"server/response"
	"server/service/actsvc"
	"server/utils/promisetool"
)

// Exit 账号退出登录
/**
header:
	Bearer Token
	username: CZKJ991706348185
*/
func Exit(context *gin.Context) {
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
	// 查询数据库，将对应的账户个人信息返回给前端
	username := claims[cm.Username]
	promise := promisetool.ToInt(claims[cm.Promise].(string))
	switch promise {
	case mc.USER, mc.ADMIN, mc.EMPLOYEE:
		err = actsvc.Exit(username.(string))
		if err != nil {
			levellog.Controller("服务器错误退出失败")
			response.Failed(context, "服务器错误退出失败")
			return
		}
		response.Success(context, username, "退出成功")
	default:
		levellog.Controller(fmt.Sprintf("权限不正确，promise: %d", promise))
		response.Failed(context, "权限不正确")
		return
	}
}
