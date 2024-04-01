package supctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	cm "server/controller/args/claims"
	mc "server/models/code"
	"server/response"
	"server/service/supsvc"
	"server/utils/promisetool"
	"strconv"
)

func Delete(context *gin.Context) {
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限校验失败")
		response.Failed(context, "权限校验失败")
		return
	}
	// 获取请求方的权限，如果是普通用户权限则无权访问
	promise := promisetool.ToInt(claims[cm.Promise].(string))
	if promise < mc.ADMIN {
		levellog.Controller("权限不够，无权访问用户列表")
		response.Failed(context, "权限不够，无权访问")
		return
	}

	id, err := strconv.Atoi(context.Query("id"))
	if err != nil {
		levellog.Controller("id转换失败")
		response.Failed(context, "id转换失败")
		return
	}
	if id <= 0 {
		levellog.Controller("id参数错误")
		response.Failed(context, "id参数错误")
		return
	}

	err = supsvc.Delete(int64(id))
	if err != nil {
		w := fmt.Sprintf("删除失败，参数错误，id = %d", id)
		levellog.Controller(w)
		response.Failed(context, w)
		return
	}
	response.Success(context, 0, "删除成功")
}
