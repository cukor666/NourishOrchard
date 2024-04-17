package fruitctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	mc "server/models/code"
	"server/response"
	"server/service/fruitsvc"
)

// Detail 水果详情
func Detail(context *gin.Context) {
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限校验失败")
		response.Failed(context, "权限校验失败")
		return
	}

	promise := controller.GetPromise(claims)
	switch promise {
	case mc.USER, mc.EMPLOYEE, mc.ADMIN:
		detailHandle(context)
	default:
		levellog.Controller("权限不足")
		response.Failed(context, "权限不足")
	}
}

func detailHandle(ctx *gin.Context) {
	type reqStruct struct {
		ID uint `form:"id" binding:"required,gt=0"`
	}

	var req reqStruct
	if err := ctx.ShouldBindQuery(&req); err != nil {
		w := fmt.Sprintf("参数校验失败，err: %s", err.Error())
		levellog.Controller(w)
		response.Failed(ctx, w)
		return
	}

	res, err := fruitsvc.Detail(int(req.ID))
	if err != nil {
		levellog.Controller("水果详情请求接口错误")
		response.Failed(ctx, "水果详情请求接口错误")
		return
	}

	response.Success(ctx, res, "查询成功")
}
