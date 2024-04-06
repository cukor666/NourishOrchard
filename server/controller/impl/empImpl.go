package impl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/models"
	"server/response"
	"server/service/empsvc"
)

type EmployeeImpl struct{}

func (ei *EmployeeImpl) Update(ctx *gin.Context, username string) {
	var emp models.Employee
	err := ctx.ShouldBindJSON(&emp)
	if err != nil {
		levellog.Controller("数据绑定失败")
		response.Failed(ctx, "数据绑定失败")
		return
	}
	if username != emp.Username {
		levellog.Controller(fmt.Sprintf("前端传递参数与JWT中不一样，拒绝请求, request: %v", emp.Username))
		response.Failed(ctx, "参数异常，拒绝请求")
		return
	}
	err = empsvc.Update(emp)
	if err != nil {
		levellog.Controller("更新失败")
		response.Failed(ctx, "更新失败")
		return
	}
	response.Success(ctx, 0, "更新成功")
}

func (ei *EmployeeImpl) Info(ctx *gin.Context, username string) {
	emp, err := empsvc.Info(username)
	if err != nil {
		response.Failed(ctx, "账号不规范")
		return
	}
	response.Success(ctx, gin.H{
		"promise": "employee",
		"data":    emp,
	}, "查询成功")
}
