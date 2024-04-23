package fruitctrl

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	mc "server/models/code"
	"server/response"
	"server/service/fruitsvc"
)

func OriginPie(context *gin.Context) {
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限校验失败")
		response.Failed(context, "权限校验失败")
		return
	}
	promise := controller.GetPromise(claims)
	switch promise {
	case mc.USER, mc.EMPLOYEE, mc.ADMIN:
		originPieHandle(context)
	default:
		levellog.Controller("权限不足")
		response.Failed(context, "权限不足")
		return
	}
}

type opRes struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

func originPieHandle(ctx *gin.Context) {
	var res []opRes
	result, err := fruitsvc.OriginPie()
	if err != nil {
		levellog.Controller("获取水果原产地饼状图失败")
		response.Failed(ctx, "获取水果原产地饼状图失败")
		return
	}
	for k, v := range result {
		res = append(res, opRes{k, v})
	}
	response.Success(ctx, res, "获取水果原产地饼状图成功")
}
