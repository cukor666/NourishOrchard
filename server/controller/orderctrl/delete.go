package orderctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"server/common/levellog"
	"server/controller"
	mc "server/models/code"
	"server/response"
	"server/service/ordersvc"
)

// Delete 删除订单接口
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
		deleteHandle(context, claims)
	default:
		levellog.Controller("无权限操作")
		response.Failed(context, "无权限操作")
	}
}

func deleteHandle(ctx *gin.Context, claims jwt.MapClaims) {
	username := controller.GetUsername(claims)
	type reqStruct struct {
		ID int64 `json:"id" form:"id"`
	}
	var req reqStruct
	if err := ctx.ShouldBind(&req); err != nil {
		response.Failed(ctx, "参数错误")
		return
	}
	err := ordersvc.Delete(req.ID, username)
	if err != nil {
		levellog.Controller("删除订单失败")
		response.Failed(ctx, "删除订单失败")
		return
	}
	response.Success(ctx, 0, "删除订单成功")
}
