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

// Insert 添加订单接口
func Insert(context *gin.Context) {
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限校验失败")
		response.Failed(context, "权限校验失败")
		return
	}
	promise := controller.GetPromise(claims)
	switch promise {
	case mc.USER, mc.ADMIN:
		insertHandle(context, claims)
	default:
		response.Failed(context, "无权限操作")
	}
}

func insertHandle(ctx *gin.Context, claims jwt.MapClaims) {
	username := controller.GetUsername(claims)
	var (
		req insertReq
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		levellog.Controller(fmt.Sprintf("参数错误，err:%s", err.Error()))
		response.Failed(ctx, "参数错误")
		return
	}
	err := ordersvc.Insert(req.toOrder(), username)
	if err != nil {
		levellog.Controller("添加订单失败")
		response.Failed(ctx, "添加订单失败")
		return
	}
	response.Success(ctx, 0, "添加订单成功")
}
