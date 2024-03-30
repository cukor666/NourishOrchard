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

// RemoveUser 彻底删除用户
/**
header:
	Bearer Token
	username: CZKJ991706348185
params:
	query:
		id: 6
		username: CZKJ18543453
*/
func RemoveUser(context *gin.Context) {
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

	type reqType struct {
		ID       uint   `json:"id" form:"id" binding:"required,gt=0"`
		Username string `json:"username" form:"username" binding:"required,username"`
	}
	var req reqType

	err = context.ShouldBindQuery(&req)
	if err != nil {
		w := spliterr.GetErrMsg(err.Error())
		levellog.Controller(fmt.Sprintf("参数校验失败，err: %s", w))
		response.Failed(context, fmt.Sprintf("参数校验失败，err: %s", w))
		return
	}

	err = lgtusersvc.RemoveUser(req.ID, req.Username)
	if err != nil {
		levellog.Controller("删除用户失败")
		response.Failed(context, "删除用户失败")
		return
	}
	response.Success(context, 0, "删除用户成功")
}
