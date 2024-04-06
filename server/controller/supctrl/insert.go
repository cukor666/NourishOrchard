package supctrl

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
	"server/service/supsvc"
	"server/utils/promisetool"
)

func Insert(context *gin.Context) {
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

	var req supReq

	err = context.ShouldBindJSON(&req)
	if err != nil {
		w := fmt.Sprintf("请求参数绑定失败，err: %s", err.Error())
		levellog.Controller(w)
		response.Failed(context, w)
		return
	}

	ok := supsvc.Insert(req.toSupplier())
	if !ok {
		levellog.Controller("参数错误，插入失败")
		response.Failed(context, "参数错误插入失败")
		return
	}
	response.Success(context, 0, "供应商信息添加成功")

}
