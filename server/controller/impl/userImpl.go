package impl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/models"
	"server/response"
	"server/service"
)

type UserImpl struct{}

func (ui *UserImpl) Update(ctx *gin.Context, username string) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		levellog.Controller("数据绑定失败")
		response.Failed(ctx, "数据绑定失败")
		return
	}
	if username != user.Username {
		levellog.Controller(fmt.Sprintf("前端传递参数与JWT中不一样，拒绝请求, request: %v", user.Username))
		response.Failed(ctx, "参数异常，拒绝请求")
		return
	}
	err = service.UserService{}.Update(user)
	if err != nil {
		levellog.Controller("更新失败")
		response.Failed(ctx, "更新失败")
		return
	}
	response.Success(ctx, 0, "更新成功")
}

func (ui *UserImpl) Info(ctx *gin.Context, username string) {
	user, err := service.UserService{}.Info(username)
	if err != nil {
		response.Failed(ctx, "账号不规范")
		return
	}
	response.Success(ctx, gin.H{
		"promise": "user",
		"data":    user,
	}, "查询成功")
}