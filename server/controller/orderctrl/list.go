package orderctrl

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	mc "server/models/code"
	"server/response"
)

// List 订单列表接口
func List(context *gin.Context) {
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限校验失败")
		response.Failed(context, "权限校验失败")
		return
	}
	promise := controller.GetPromise(claims)
	switch promise {
	case mc.USER:
		listHandleUser(context, claims)
	case mc.Employed, mc.ADMIN:
		listHandle(context)
	default:
		levellog.Controller("权限不足")
		response.Failed(context, "权限不足")
	}
}
