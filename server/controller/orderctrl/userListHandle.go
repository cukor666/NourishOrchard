package orderctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"server/common/levellog"
	"server/common/simpletool"
	"server/controller"
	"server/response"
	"server/service/ordersvc"
)

func listHandleUser(ctx *gin.Context, claims jwt.MapClaims) {
	var (
		p simpletool.Page
	)
	if err := ctx.ShouldBindQuery(&p); err != nil {
		levellog.Controller("参数错误")
		response.Failed(ctx, "参数错误")
		return
	}
	username := controller.GetUsername(claims)
	orders, total, err := ordersvc.ListForUser(p, username)
	if err != nil {
		levellog.Controller("获取订单列表失败")
		response.Failed(ctx, "获取订单列表失败")
		return
	}
	response.Success(ctx, gin.H{
		"orders": orders,
		"total":  total,
	}, "获取订单列表成功")
}
