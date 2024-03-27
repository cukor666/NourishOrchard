package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/common/simpletool"
	"server/config"
	cm "server/controller/args/claims"
	"server/controller/args/header"
	"server/models"
	mc "server/models/code"
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
		pageNum: 1
		id: 3
		username: CZKJ991706344513
		... user的字段信息

如果user字段的值是零值，则服务器端不做该字段的SQL拼接

http://localhost:9000/user/list?pageSize=3&pageNum=2&id=3&username=CZKJ991706344513 ...
*/
func (uc *UserController) List(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levellog.Controller("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		levellog.Controller("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	// 获取请求方的权限，如果是普通用户权限则无权访问
	promise := utils.PromiseToInt(claims[cm.Promise].(string))
	if promise <= mc.USER {
		levellog.Controller("权限不够，无权访问用户列表")
		response.Failed(context, "权限不够，无权访问")
		return
	}

	// 权限校验通过之后获取参数
	var (
		p    simpletool.Page
		user models.User
	)

	// 获取分页的参数
	pageSize := context.Query("pageSize")
	if p.Size, err = strconv.Atoi(pageSize); err != nil {
		levellog.Controller(fmt.Sprintf("pageSize参数错误, pageSize: %v", p.Size))
		response.Failed(context, "pageSize参数错误")
		return
	}
	pageNum := context.Query("pageNum")
	if p.Num, err = strconv.Atoi(pageNum); err != nil {
		levellog.Controller(fmt.Sprintf("pageNum参数错误, pageNum参数错误: %v", p.Num))
		response.Failed(context, "pageNum参数错误")
		return
	}

	// 获取查询用户的参数
	err = context.ShouldBindQuery(&user)
	if err != nil {
		levellog.Controller(fmt.Sprintf("获取用户参数失败, user: %v", user))
		response.Failed(context, "获取用户相关参数失败")
		return
	}

	users, total, err := service.UserService{}.List(p, user)
	if err != nil {
		levellog.Controller("服务端错误，查询失败")
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
func (uc *UserController) Update(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levellog.Controller("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		levellog.Controller("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	// 获取请求方的权限，如果是普通用户权限则无权访问
	promise := utils.PromiseToInt(claims[cm.Promise].(string))
	if promise < mc.ADMIN {
		levellog.Controller("权限不够，更新用户信息")
		response.Failed(context, "权限不够，无权操作")
		return
	}
	var (
		user models.User
	)
	err = context.ShouldBindJSON(&user)
	if err != nil {
		levellog.Controller("前端数据绑定失败")
		response.Failed(context, "数据绑定失败")
		return
	}
	levellog.Controller(fmt.Sprintf("user: %v", user))
	err = service.UserService{}.Update(user)
	if err != nil {
		levellog.Controller("更新失败")
		response.Failed(context, "系统错误更新失败")
		return
	}
	response.Success(context, user, "更新用户信息成功")
}

// Delete 删除用户接口
func (uc *UserController) Delete(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levellog.Controller("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		levellog.Controller("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	// 获取请求方的权限，如果是普通用户权限则无权访问
	promise := utils.PromiseToInt(claims[cm.Promise].(string))
	if promise < mc.ADMIN {
		levellog.Controller("权限不够，更新用户信息")
		response.Failed(context, "权限不够，无权操作")
		return
	}

	username := context.Query("username")
	if username == "" {
		levellog.Controller("参数错误")
		response.Failed(context, "参数错误")
		return
	}
	ok := valid.Username(username)
	if !ok {
		levellog.Controller("账号校验失败")
		response.Failed(context, "参数校验失败")
		return
	}
	deleteUser, err := service.UserService{}.DeleteUser(username)
	if err != nil {
		levellog.Controller(fmt.Sprintf("删除用户信息时出错，%v", err))
		response.Failed(context, "删除失败")
		return
	}
	response.Success(context, deleteUser, "删除成功")
}

// LogoutList 查询注销用户列表 要求分页查询，并将列表按id倒序
/**
header:
	Bearer Token
	username: CZKJ991706348185
params:
	query:
		pageSize: 3
		pageNum: 2

http://localhost:9000/user/logout-list?pageSize=3&pageNum=2
*/
func (uc *UserController) LogoutList(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levellog.Controller("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		levellog.Controller("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	// 获取请求方的权限，如果是普通用户权限则无权访问
	promise := utils.PromiseToInt(claims[cm.Promise].(string))
	if promise <= mc.USER {
		levellog.Controller("权限不够，无权访问用户列表")
		response.Failed(context, "权限不够，无权访问")
		return
	}

	// 权限校验通过之后获取参数
	var (
		p          simpletool.Page
		logoutUser models.LogoutUser
	)

	pageSize := context.Query("pageSize")
	if p.Size, err = strconv.Atoi(pageSize); err != nil {
		levellog.Controller(fmt.Sprintf("pageSize参数错误, pageSize: %v", p.Size))
		response.Failed(context, "pageSize参数错误")
		return
	}
	pageNum := context.Query("pageNum")
	if p.Num, err = strconv.Atoi(pageNum); err != nil {
		levellog.Controller(fmt.Sprintf("pageNum参数错误, pageNum参数错误: %v", p.Num))
		response.Failed(context, "pageNum参数错误")
		return
	}

	err = context.ShouldBindQuery(&logoutUser)
	if err != nil {
		levellog.Controller(fmt.Sprintf("获取注销用户参数失败, logoutUser: %v", logoutUser))
		response.Failed(context, "获取注销用户参数失败")
		return
	}

	users, total, err := service.UserService{}.LogoutList(p, logoutUser)
	if err != nil {
		levellog.Controller("服务端错误，查询失败")
		response.Failed(context, "服务端错误，查询失败")
		return
	}
	response.Success(context, gin.H{
		"users": users,
		"total": total,
	}, "查询注销用户列表成功")
}

// RecoverUser 恢复用户信息
/**
header:
	Bearer Token
	username: CZKJ991706348185
body:
	{
		"username": "CZKJ991706340782"
	}
*/
func (uc *UserController) RecoverUser(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levellog.Controller("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		levellog.Controller("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	// 获取请求方的权限，如果是普通用户权限则无权访问
	promise := utils.PromiseToInt(claims[cm.Promise].(string))
	if promise < mc.ADMIN {
		levellog.Controller("权限不够，更新用户信息")
		response.Failed(context, "权限不够，无权操作")
		return
	}

	type tempRequest struct {
		Username string `json:"username"`
	}
	var req tempRequest
	err = context.ShouldBindJSON(&req)
	if err != nil {
		levellog.Controller(fmt.Sprintf("参数绑定失败，req = %v", req))
		response.Failed(context, "参数绑定失败")
		return
	}
	ok := valid.Username(req.Username)
	if !ok {
		levellog.Controller("账号校验失败")
		response.Failed(context, "参数校验失败")
		return
	}
	user, err := service.UserService{}.RecoverUser(req.Username)
	if err != nil {
		levellog.Controller(fmt.Sprintf("恢复用户信息失败，user = %v", user))
		response.Failed(context, "恢复用户信息失败")
		return
	}
	response.Success(context, user, "恢复用户信息成功")
}

// RemoveUser 彻底删除用户
/**
header:
	Bearer Token
	username: CZKJ991706348185
params:
	query:
		id: 6
		username: CZKJ18543453
*/
func (uc *UserController) RemoveUser(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levellog.Controller("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		levellog.Controller("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	// 获取请求方的权限，如果是普通用户权限则无权访问
	promise := utils.PromiseToInt(claims[cm.Promise].(string))
	if promise < mc.ADMIN {
		levellog.Controller("权限不够，无权访问用户列表")
		response.Failed(context, "权限不够，无权访问")
		return
	}

	type reqType struct {
		ID       uint   `json:"id" form:"id"`
		Username string `json:"username" form:"username"`
	}
	var req reqType

	err = context.ShouldBindQuery(&req)
	if err != nil {
		levellog.Controller(fmt.Sprintf("参数绑定失败， req = %v", req))
		response.Failed(context, "参数绑定失败")
		return
	}

	err = service.UserService{}.RemoveUser(req.ID, req.Username)
	if err != nil {
		levellog.Controller("删除用户失败")
		response.Failed(context, "删除用户失败")
		return
	}
	response.Success(context, 0, "删除用户成功")
}
