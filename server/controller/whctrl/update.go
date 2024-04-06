package whctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	mc "server/models/code"
	"server/response"
	"server/service/whsvc"
)

// Update 更新仓库信息接口
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
		updateHandle(context)
	default:
		levellog.Controller("权限不够，拒绝请求")
		response.Failed(context, "权限不够，拒绝请求")
	}
}

func updateHandle(ctx *gin.Context) {
	var req updateReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		w := fmt.Sprintf("参数校验失败，err: %s", err.Error())
		levellog.Controller(w)
		response.Failed(ctx, w)
		return
	}

	err = whsvc.Update(req.toWarehouse())
	if err != nil {
		w := fmt.Sprintf("更新失败，可能是数据库中不存在id为%d的仓库数据", req.ID)
		levellog.Controller(w)
		response.Failed(ctx, w)
		return
	}
	response.Success(ctx, 0, "更新成功")
}
