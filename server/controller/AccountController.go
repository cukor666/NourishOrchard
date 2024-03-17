package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/config"
	cm "server/controller/args/claims"
	"server/controller/args/header"
	"server/models"
	mc "server/models/code"
	"server/request"
	"server/response"
	"server/service"
	"server/utils"
)

/**
AccountController是关于个人管理的接口控制管理，对应的大批管理由对应的Controller管理
例如：用户列表由UserController自行管理
*/

// GetAccount 获取账户信息
/**
header:
	Bearer Token
	username: CZKJ991706348185
body: 空
*/
func (a AccountController) GetAccount(context *gin.Context) {
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
	// 查询数据库，将对应的账户个人信息返回给前端
	username := claims[cm.Username]
	promise := utils.PromiseToInt(claims[cm.Promise].(string))
	switch promise {
	case mc.USER:
		// 调用UserService
		user, err := service.UserService{}.Info(username.(string))
		if err != nil {
			response.Failed(context, "账号不规范")
			return
		}
		response.Success(context, gin.H{
			"promise": "user",
			"data":    user,
		}, "查询成功")
	case mc.EMPLOYEE:
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
	case mc.ADMIN:
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

// Update 更新用户个人信息
/**
header:
	Bearer Token
	username: CZKJ991706348185
body:
	{
		"username": "CZKJ991706348185",
		"name": "Coco大美女",
		"phone": "13489427501",
		"position": "仓库管理员",
		"salary": 3000
	}
*/
func (a AccountController) Update(context *gin.Context) {
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
	// 查询数据库，将对应的账户个人信息返回给前端
	username := claims[cm.Username]
	promise := utils.PromiseToInt(claims[cm.Promise].(string))
	var (
		user     models.User
		admin    models.Admin
		employee models.Employee
	)
	switch promise {
	case mc.USER:
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
		err = service.UserService{}.Update(user)
	case mc.EMPLOYEE:
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
	case mc.ADMIN:
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

// Exit 账号退出登录
/**
header:
	Bearer Token
	username: CZKJ991706348185
*/
func (a AccountController) Exit(context *gin.Context) {
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
	// 查询数据库，将对应的账户个人信息返回给前端
	username := claims[cm.Username]
	promise := utils.PromiseToInt(claims[cm.Promise].(string))
	switch promise {
	case mc.USER, mc.ADMIN, mc.EMPLOYEE:
		err = service.AccountService{}.Exit(username.(string))
		if err != nil {
			levelLog("服务器错误退出失败")
			response.Failed(context, "服务器错误退出失败")
			return
		}
		response.Success(context, username, "退出成功")
	default:
		levelLog(fmt.Sprintf("权限不正确，promise: %d", promise))
		response.Failed(context, "权限不正确")
		return
	}
}

// ChangePassword 修改账户密码
func (a AccountController) ChangePassword(context *gin.Context) {
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
	username := claims[cm.Username]
	promise := utils.PromiseToInt(claims[cm.Promise].(string))
	type pwdType struct {
		OldPassword string `json:"oldPassword" form:"oldPassword"`
		NewPassword string `json:"newPassword" form:"newPassword"`
	}
	var pwd pwdType
	err = context.ShouldBindJSON(&pwd)
	if err != nil {
		levelLog(fmt.Sprintf("绑定前端数据失败, pwd = %v", pwd))
		response.Failed(context, "参数错误")
		return
	}
	switch promise {
	case mc.USER, mc.ADMIN, mc.EMPLOYEE:
		err = service.AccountService{}.ChangePassword(username.(string), pwd.OldPassword, pwd.NewPassword)
		if err != nil {
			levelLog("服务器错误修改密码失败")
			response.Failed(context, "服务器错误修改密码失败")
			return
		}
		response.Success(context, username, "修改密码成功")
	default:
		levelLog(fmt.Sprintf("权限不正确，promise: %d", promise))
		response.Failed(context, "权限不正确")
		return
	}
}

// ForgetPassword 用户忘记密码
func (a AccountController) ForgetPassword(context *gin.Context) {
	var (
		req request.ForgetPwdReq
		err error
	)

	err = context.ShouldBindJSON(&req)
	if err != nil {
		levelLog(fmt.Sprintf("数据绑定失败， req = %v", req))
		response.Failed(context, "参数错误")
		return
	}
	promise := utils.PromiseToInt(req.Promise)
	if promise != mc.USER {
		levelLog(fmt.Sprintf("权限不正确，promise = %s", req.Promise))
		response.Failed(context, "权限不正确，拒绝请求")
		return
	}

	// 校验短信验证码
	myCode := "1024"
	if req.Code != myCode {
		levelLog(fmt.Sprintf("验证码错误应该得到：%s，却得到%s", myCode, req.Code))
		response.Failed(context, "验证码错误")
		return
	}

	err = service.AccountService{}.ForgetPassword(req)
	if err != nil {
		levelLog("服务器错误，忘记密码接口拒绝请求")
		response.Failed(context, "服务器错误")
		return
	}
	response.Success(context, 0, "忘记密码更改密码成功")
}
