package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common"
	"server/config"
	"server/models"
	"server/response"
	"server/service"
	"server/utils"
)

// GetAccount 获取账户信息
func (a AccountController) GetAccount(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader("Authorization")
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levelLog("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		levelLog("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	// 查询数据库，将对应的账户个人信息返回给前端
	username := claims["username"]
	promise := utils.PromiseToInt(claims["promise"].(string))
	switch promise {
	case common.USER:
		// 调用UserService
		user, err := userService.Info(username.(string))
		if err != nil {
			response.Failed(context, "账号不规范")
			return
		}
		response.Success(context, gin.H{
			"promise": "user",
			"data":    user,
		}, "查询成功")
	case common.EMPLOYEE:
		// 调用employeeService
		emp, err := service.EmployeeService{}.Info(username.(string))
		if err != nil {
			response.Failed(context, "账号不规范")
			return
		}
		response.Success(context, gin.H{
			"promise": "employee",
			"data":    emp,
		}, "查询成功")
	case common.ADMIN:
		// 调用adminService
		admin, err := service.AdminService{}.Info(username.(string))
		if err != nil {
			response.Failed(context, "账号不规范")
			return
		}
		response.Success(context, gin.H{
			"promise": "admin",
			"data":    admin,
		}, "查询成功")
	default:
		levelLog("权限不对")
		response.Failed(context, "权限不对")
	}
}

func (a AccountController) Update(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader("Authorization")
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levelLog("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		levelLog("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	// 查询数据库，将对应的账户个人信息返回给前端
	username := claims["username"]
	promise := utils.PromiseToInt(claims["promise"].(string))
	var (
		user     models.User
		admin    models.Admin
		employee models.Employee
	)
	switch promise {
	case common.USER:
		err = context.ShouldBindJSON(&user)
		if err != nil {
			levelLog("数据绑定失败")
			response.Failed(context, "数据绑定失败")
			return
		}
		if username != user.Username {
			levelLog(fmt.Sprintf("前端传递参数与JWT中不一样，拒绝请求, request: %v", user.Username))
			response.Failed(context, "参数异常，拒绝请求")
			return
		}
		// 调用UserService
		err = userService.Update(user)
	case common.EMPLOYEE:
		err = context.ShouldBindJSON(&employee)
		if err != nil {
			levelLog(fmt.Sprintf("数据绑定失败, emp: %v", employee))
			response.Failed(context, "数据绑定失败")
			return
		}
		if username != employee.Username {
			levelLog(fmt.Sprintf("前端传递参数与JWT中不一样，拒绝请求, request: %v", user.Username))
			response.Failed(context, "参数异常，拒绝请求")
			return
		}
		// 调用employeeService
		err = service.EmployeeService{}.Update(employee)
	case common.ADMIN:
		err = context.ShouldBindJSON(&admin)
		if err != nil {
			levelLog("数据绑定失败")
			response.Failed(context, "数据绑定失败")
			return
		}
		if username != admin.Username {
			levelLog(fmt.Sprintf("前端传递参数与JWT中不一样，拒绝请求, request: %v", admin.Username))
			response.Failed(context, "参数异常，拒绝请求")
			return
		}
		// 调用adminService
		err = service.AdminService{}.Update(admin)
	default:
		levelLog("权限不对")
		response.Failed(context, "权限不对")
		return
	}
	if err != nil {
		response.Failed(context, "更新失败")
		return
	}
	response.Success(context, 0, "更新成功")
}
