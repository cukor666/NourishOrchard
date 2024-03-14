package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/simpletool"
	"server/config"
	cm "server/controller/args/claims"
	"server/controller/args/header"
	"server/models"
	mc "server/models/code"
	"server/response"
	"server/service"
	"server/utils"
	"strconv"
)

// List 查询员工列表
/**
header:
	Bearer Token
	username: CZKJ991706348185
params:
	query:
		pageSize: 3
		pageNum: 2
*/
func (e EmployeeController) List(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
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
	// 获取请求方的权限，如果是普通用户权限则无权访问
	promise := utils.PromiseToInt(claims[cm.Promise].(string))
	if promise < mc.ADMIN {
		levelLog("权限不够，无权访问用户列表")
		response.Failed(context, "权限不够，无权访问")
		return
	}
	var (
		p        simpletool.Page
		employee models.Employee
	)

	pageSize := context.Query("pageSize")
	if p.Size, err = strconv.Atoi(pageSize); err != nil {
		levelLog(fmt.Sprintf("pageSize参数错误，pageSize: %v", p.Size))
		response.Failed(context, "pageSize参数错误")
		return
	}
	pageNum := context.Query("pageNum")
	if p.Num, err = strconv.Atoi(pageNum); err != nil {
		levelLog(fmt.Sprintf("pageNum参数错误，pageNum: %v", p.Num))
		response.Failed(context, "pageNum参数错误")
		return
	}

	err = context.ShouldBindQuery(&employee)
	if err != nil {
		levelLog(fmt.Sprintf("参数绑定失败, empcode = %v", employee))
		response.Failed(context, "参数绑定失败")
		return
	}

	employees, total, err := service.EmployeeService{}.ListWithPage(p, employee)
	if err != nil {
		levelLog("服务器端错误，查询失败")
		response.Failed(context, "服务器端错误，查询失败")
		return
	}
	response.Success(context, gin.H{
		"employees": employees,
		"total":     total,
	}, "查询员工列表成功")
}

// Update 更新员工信息
/**
header:
	Bearer Token
	username: CZKJ991706348185
body:
	{
		"id": 7,
		"username": "CZKJ201709005881",
		"name": "吴宣仪大美女",
		"phone": "18577659826",
		"position": "仓库管理员",
		"salary": 4000
	}
*/
func (e EmployeeController) Update(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
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
	// 获取请求方的权限，如果是普通用户权限则无权访问
	promise := utils.PromiseToInt(claims[cm.Promise].(string))
	if promise < mc.ADMIN {
		levelLog("权限不够，更新用户信息")
		response.Failed(context, "权限不够，无权操作")
		return
	}
	var (
		employee models.Employee
	)
	err = context.ShouldBindJSON(&employee)
	if err != nil {
		levelLog("前端参数绑定失败")
		response.Failed(context, "参数绑定失败")
		return
	}
	levelLog(fmt.Sprintf("employee: %v", employee))
	err = service.EmployeeService{}.Update(employee)
	if err != nil {
		levelLog("更新失败")
		response.Failed(context, "系统错误更新失败")
		return
	}
	response.Success(context, employee, "更新员工成功")
}
