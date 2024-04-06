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

// Insert 添加新仓库接口
func Insert(context *gin.Context) {
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限校验失败")
		response.Failed(context, "权限校验失败")
		return
	}
	promise := controller.GetPromise(claims)
	switch promise {
	case mc.ADMIN:
		insertHandle(context)
	default:
		levellog.Controller("权限不够")
		response.Failed(context, "权限不够，拒绝请求")
	}
}

func insertHandle(ctx *gin.Context) {
	var (
		req insertReq
	)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		levellog.Controller(fmt.Sprintf("参数req校验不通过，err: %s", err.Error()))
		response.Failed(ctx, "参数错误，请检查")
		return
	}

	err = whsvc.Insert(req.toWarehouse())
	if err != nil {
		w := fmt.Sprintf("插入失败，err：%v", err)
		levellog.Controller(w)
		response.Failed(ctx, w)
		return
	}
	response.Success(ctx, 0, "插入成功")
}
