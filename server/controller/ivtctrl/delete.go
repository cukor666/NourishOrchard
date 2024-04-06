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

// Delete 删除库存信息接口
func Delete(context *gin.Context) {
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限校验失败")
		response.Failed(context, "权限校验失败")
		return
	}
	promise := controller.GetPromise(claims)
	switch promise {
	case mc.ADMIN:
		deleteHandle(context)
	default:
		levellog.Controller("权限不够，拒绝请求")
		response.Failed(context, "权限不够拒绝请求，请检查")
	}
}

func deleteHandle(ctx *gin.Context) {
	type deleteReq struct {
		ID int64 `form:"id" binding:"required,gt=0"`
	}
	var req deleteReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		w := fmt.Sprintf("参数校验失败，err: %s", err.Error())
		levellog.Controller(w)
		response.Failed(ctx, w)
		return
	}
	err = ivtsvc.Delete(req.ID)
	if err != nil {
		w := fmt.Sprintf("删除失败，err: %s", err.Error())
		levellog.Controller(w)
		response.Failed(ctx, w)
		return
	}
	response.Success(ctx, 0, "删除成功")
}
