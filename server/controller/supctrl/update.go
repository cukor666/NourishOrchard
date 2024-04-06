package supctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	mc "server/models/code"
	"server/response"
	"server/service/supsvc"
)

func Update(context *gin.Context) {
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限校验失败")
		response.Failed(context, "权限校验失败")
		return
	}
	promise := controller.GetPromise(claims)
	switch promise {
	case mc.ADMIN:
		handle(context)
	default:
		levellog.Controller("权限不正确，拒绝请求")
		response.Failed(context, "权限不正确拒绝请求")
	}
}

func handle(context *gin.Context) {
	var req supUpdateReq
	err := context.ShouldBindJSON(&req)
	if err != nil {
		w := fmt.Sprintf("参数校验失败，err: %s", err.Error())
		levellog.Controller(w)
		response.Failed(context, w)
		return
	}
	levellog.Controller("数据校验成功")
	err = supsvc.Update(req.toSupplier())
	if err != nil {
		levellog.Controller("更新供应商信息失败")
		response.Failed(context, "更新供应商信息失败")
		return
	}
	response.Success(context, req, "更新供应商信息成功")
}
