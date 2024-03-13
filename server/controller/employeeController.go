package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common"
	"server/common/simpletool"
	"server/config"
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
	// 获取请求方的权限，如果是普通用户权限则无权访问
	promise := utils.PromiseToInt(claims["promise"].(string))
	if promise < common.ADMIN {
		levelLog("权限不够，无权访问用户列表")
		response.Failed(context, "权限不够，无权访问")
		return
	}
	var p simpletool.Page
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
	employees, total, err := service.EmployeeService{}.ListWithPage(p)
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