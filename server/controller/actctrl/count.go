package actctrl

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller"
	"server/dto/actdto"
	mc "server/models/code"
	"server/response"
	"server/service/actsvc"
)

func Count(context *gin.Context) {
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限校验失败")
		response.Failed(context, "权限校验失败")
		return
	}
	promise := controller.GetPromise(claims)
	switch promise {
	case mc.USER, mc.EMPLOYEE, mc.ADMIN:
		countHandle(context)
	default:
		levellog.Controller("权限不足")
		response.Failed(context, "权限不足")
	}
}

type cntRes struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func (c *cntRes) set(obj actdto.Count) {
	switch obj.Promise {
	case mc.USER:
		c.Name = "用户"
	case mc.EMPLOYEE:
		c.Name = "员工"
	case mc.ADMIN:
		c.Name = "管理员"
	default:
		c.Name = "未知"
	}
	c.Value = obj.Cnt
}

func countHandle(ctx *gin.Context) {
	var res []cntRes
	result, err := actsvc.Count()
	if err != nil {
		levellog.Controller("获取人员数量失败")
		response.Failed(ctx, "获取人员数量失败")
		return
	}

	for _, r := range result {
		var c cntRes
		c.set(r)
		res = append(res, c)
	}

	response.Success(ctx, res, "获取人员数量成功")
}
