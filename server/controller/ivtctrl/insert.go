package ivtctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	mc "server/models/code"
	"server/response"
	"server/service/ivtsvc"
)

// Insert 添加库存信息接口
func Insert(context *gin.Context) {
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限校验失败")
		response.Failed(context, "权限校验失败")
		return
	}
	promise := controller.GetPromise(claims)
	switch promise {
	case mc.EMPLOYEE, mc.ADMIN:
		insertHandle(context)
	default:
		levellog.Controller("权限不够")
		response.Failed(context, "权限不够，拒绝请求")
	}
}

func insertHandle(ctx *gin.Context) {
	var req insertReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		w := fmt.Sprintf("参数校验失败，err: %s", err.Error())
		levellog.Controller(w)
		response.Failed(ctx, w)
		return
	}

	if err := ivtsvc.Insert(req.toInventory()); err != nil {
		w := fmt.Sprintf("添加失败，err: %s", err.Error())
		levellog.Controller(w)
		response.Failed(ctx, w)
		return
	}
	response.Success(ctx, 0, "添加成功")
}
