package impl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/models"
	"server/response"
	"server/service/adminsvc"
)

type AdminImpl struct{}

func (a *AdminImpl) Update(ctx *gin.Context, username string) {
	var admin models.Admin
	err := ctx.ShouldBindJSON(&admin)
	if err != nil {
		levellog.Controller("数据绑定失败")
		response.Failed(ctx, "数据绑定失败")
		return
	}
	if username != admin.Username {
		levellog.Controller(fmt.Sprintf("前端传递参数与JWT中不一样，拒绝请求, request: %v", admin.Username))
		response.Failed(ctx, "参数异常，拒绝请求")
		return
	}
	err = adminsvc.Update(admin)
	if err != nil {
		levellog.Controller("更新失败")
		response.Failed(ctx, "更新失败")
		return
	}
	response.Success(ctx, 0, "更新成功")
}

func (a *AdminImpl) Info(ctx *gin.Context, username string) {
	admin, err := adminsvc.Info(username)
	if err != nil {
		response.Failed(ctx, "账号不规范")
		return
	}
	response.Success(ctx, gin.H{
		"promise": "admin",
		"data":    admin,
	}, "查询成功")
}
