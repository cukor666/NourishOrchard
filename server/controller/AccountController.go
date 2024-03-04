package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/common"
	"server/config"
	"server/models"
	"server/response"
	"server/utils"
)

func (a AccountController) GetAccount(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader("Authorization")
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		log.Println("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		log.Println("解析token失败")
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
			response.Failed(context, "用户名不规范")
			return
		}
		response.Success(context, user, "查询成功")
	case common.EMPLOYEE:
	// 调用employeeService
	case common.ADMIN:
	// 调用adminService
	default:
		log.Println("权限不对")
		response.Failed(context, "权限不对")
	}
}

func (a AccountController) Update(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader("Authorization")
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		log.Println("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		log.Println("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	// 查询数据库，将对应的账户个人信息返回给前端
	username := claims["username"]
	promise := utils.PromiseToInt(claims["promise"].(string))
	var (
		user models.User
	)
	err = context.ShouldBindJSON(&user)
	if err != nil {
		log.Println("数据绑定失败")
		response.Failed(context, "数据绑定失败")
		return
	}
	switch promise {
	case common.USER:
		if username != user.Username {
			log.Println("前端传递参数与JWT中不一样，拒绝请求")
			response.Failed(context, "参数异常，拒绝请求")
			return
		}
		// 调用UserService
		err = userService.Update(user)
		if err != nil {
			response.Failed(context, "更新失败")
			return
		}
		response.Success(context, 0, "更新成功")
	case common.EMPLOYEE:
	// 调用employeeService
	case common.ADMIN:
	// 调用adminService
	default:
		log.Println("权限不对")
		response.Failed(context, "权限不对")
	}
}
