package userctrl

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
	"server/models"
	mc "server/models/code"
	"server/response"
	"server/service/usersvc"
	"server/utils/promisetool"
)

// List 查询用户列表 要求分页查询，并将列表按id倒序
/**
header:
	Bearer Token
	username: CZKJ991706348185
params:
	query:
		pageSize: 3
		pageNum: 1
		id: 3
		username: CZKJ991706344513
		... user的字段信息

如果user字段的值是零值，则服务器端不做该字段的SQL拼接

http://localhost:9000/user/list?pageSize=3&pageNum=2&id=3&username=CZKJ991706344513 ...
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
	if promise <= mc.USER {
		levellog.Controller("权限不够，无权访问用户列表")
		response.Failed(context, "权限不够，无权访问")
		return
	}

	// 权限校验通过之后获取参数
	var (
		p    simpletool.Page
		user models.User
	)

	err = context.ShouldBindQuery(&p)
	if err != nil {
		msg := spliterr.GetErrMsg(err.Error())
		w := fmt.Sprintf("分页参数校验失败, err: %s", msg)
		levellog.Controller(w)
		response.Failed(context, w)
		return
	}

	// 获取查询用户的参数
	err = context.ShouldBindQuery(&user)
	if err != nil {
		msg := spliterr.GetErrMsg(err.Error())
		w := fmt.Sprintf("获取用户参数失败, err: %s", msg)
		levellog.Controller(w)
		response.Failed(context, w)
		return
	}

	users, total, err := usersvc.List(p, user)
	if err != nil {
		levellog.Controller("服务端错误，查询失败")
		response.Failed(context, "服务端错误，查询失败")
		return
	}
	response.Success(context, gin.H{
		"users": users,
		"total": total,
	}, "查询用户列表成功")
}
