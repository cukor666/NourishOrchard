package orderctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"server/common/levellog"
	"server/controller"
	mc "server/models/code"
	"server/response"
	"server/service/ordersvc"
)

// Update 更新订单接口
func Update(context *gin.Context) {
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限验证失败")
		response.Failed(context, "权限验证失败")
		return
	}
	promise := controller.GetPromise(claims)
	switch promise {
	case mc.ADMIN:
		updateHandle(context, claims)
	default:
		levellog.Controller("权限不足")
		response.Failed(context, "权限不足")
	}
}

func updateHandle(ctx *gin.Context, claims jwt.MapClaims) {
	var (
		req updateReq
	)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		levellog.Controller(fmt.Sprintf("参数错误，msg:%s", err.Error()))
		response.Failed(ctx, "参数错误")
		return
	}
	username := controller.GetUsername(claims)
	err = ordersvc.Update(req.toOrder(), username)
	if err != nil {
		levellog.Controller("更新订单失败")
		response.Failed(ctx, "更新订单失败")
		return
	}
	response.Success(ctx, 0, "更新订单成功")
}
