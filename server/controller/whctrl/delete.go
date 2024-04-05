package whctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	mc "server/models/code"
	"server/response"
	"server/service/whsvc"
	"strconv"
)

// Delete 删除仓库信息接口
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
		response.Failed(context, "权限不够，拒绝请求")
	}
}

func deleteHandle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		levellog.Controller("获取参数id失败")
		response.Failed(ctx, "参数错误，拒绝请求，请检查参数")
		return
	}

	ok := validId(id)
	if !ok {
		levellog.Controller("参数id校验失败")
		response.Failed(ctx, "参数校验失败，请重试")
		return
	}

	err = whsvc.Delete(int64(id))
	if err != nil {
		levellog.Controller("系统错误，删除失败")
		response.Failed(ctx, fmt.Sprintf("系统错误，删除失败，或者数据库中无id为%d的仓库", id))
		return
	}
	response.Success(ctx, 0, "删除成功")
}

func validId(id int) bool {
	if id <= 0 {
		return false
	}
	return true
}
