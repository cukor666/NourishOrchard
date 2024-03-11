package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common"
	"server/common/simpletool"
	"server/config"
	"server/models"
	"server/response"
	"server/service"
	"server/utils"
	"server/utils/valid"
	"strconv"
)

// List 查询用户列表 要求分页查询，并将列表按id倒序
/**
header:
	Bearer Token
	username: CZKJ991706348185
params:
	query:
		pageSize: 3
		pageNum: 2

http://localhost:9000/user/list?pageSize=3&pageNum=2
*/
func (u UserController) List(context *gin.Context) {
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
	if promise <= common.USER {
		levelLog("权限不够，无权访问用户列表")
		response.Failed(context, "权限不够，无权访问")
		return
	}

	// 权限校验通过之后获取参数
	var p simpletool.Page
	pageSize := context.Query("pageSize")
	if p.Size, err = strconv.Atoi(pageSize); err != nil {
		levelLog(fmt.Sprintf("pageSize参数错误, pageSize: %v", p.Size))
		response.Failed(context, "pageSize参数错误")
		return
	}
	pageNum := context.Query("pageNum")
	if p.Num, err = strconv.Atoi(pageNum); err != nil {
		levelLog(fmt.Sprintf("pageNum参数错误, pageNum参数错误: %v", p.Num))
		response.Failed(context, "pageNum参数错误")
		return
	}

	users, total, err := userService.List(p)
	if err != nil {
		levelLog("服务端错误，查询失败")
		response.Failed(context, "服务端错误，查询失败")
		return
	}
	response.Success(context, gin.H{
		"users": users,
		"total": total,
	}, "查询用户列表成功")
}

// Update 更新用户信息
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
		"gender": "男",
		"address": "海南省海口市",
		"birthday": "1995-03-16T00:00:00Z"
	}
*/
func (u UserController) Update(context *gin.Context) {
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
		levelLog("权限不够，更新用户信息")
		response.Failed(context, "权限不够，无权操作")
		return
	}
	var (
		user models.User
	)
	err = context.ShouldBindJSON(&user)
	if err != nil {
		levelLog("前端数据绑定失败")
		response.Failed(context, "数据绑定失败")
		return
	}
	levelLog(fmt.Sprintf("user: %v", user))
	err = service.UserService{}.Update(user)
	if err != nil {
		levelLog("更新失败")
		response.Failed(context, "系统错误更新失败")
		return
	}
	response.Success(context, user, "更新用户信息成功")
}

// Delete 删除用户接口
func (u UserController) Delete(context *gin.Context) {
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
		levelLog("权限不够，更新用户信息")
		response.Failed(context, "权限不够，无权操作")
		return
	}

	username := context.Query("username")
	if username == "" {
		levelLog("参数错误")
		response.Failed(context, "参数错误")
		return
	}
	ok := valid.ValidUsername(username)
	if !ok {
		levelLog("账号校验失败")
		response.Failed(context, "参数校验失败")
		return
	}
	deleteUser, err := service.UserService{}.DeleteUser(username)
	if err != nil {
		levelLog(fmt.Sprintf("删除用户信息时出错，%v", err))
		response.Failed(context, "删除失败")
		return
	}
	response.Success(context, deleteUser, "删除成功")
}
