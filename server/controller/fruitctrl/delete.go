package fruitctrl

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
	"server/service/fruitsvc"
	"server/utils/promisetool"
	"strconv"
)

// Delete 删除水果信息
func Delete(context *gin.Context) {
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
	promise := promisetool.ToInt(claims[cm.Promise].(string))
	switch promise {
	case mc.USER:
		levellog.Controller("权限不够")
		response.Failed(context, "权限不够")
	case mc.EMPLOYEE, mc.ADMIN:
		id, err := strconv.Atoi(context.Query("id"))
		if err != nil {
			levellog.Controller("获取id失败")
			response.Failed(context, "参数错误")
			return
		}
		if id <= 0 {
			w := fmt.Sprintf("id参数校验不通过，id = %d", id)
			levellog.Valid(w)
			response.Failed(context, w)
			return
		}
		fruit, err := fruitsvc.Delete(id)
		if err != nil {
			levellog.Controller("删除失败")
			response.Failed(context, "系统错误删除失败")
			return
		}
		response.Success(context, fruit, "删除成功")
	default:
		levellog.Controller("未知权限")
		response.Failed(context, "未知权限")
	}
}
