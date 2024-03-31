package lgtuserctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/common/simpletool"
	"server/config"
	"server/controller"
	cm "server/controller/args/claims"
	"server/controller/args/header"
	"server/models"
	mc "server/models/code"
	"server/response"
	"server/service/lgtusersvc"
	"server/utils/promisetool"
	"sync"
)

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
func LogoutList(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := controller.GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
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
	promise := promisetool.ToInt(claims[cm.Promise].(string))
	if promise <= mc.USER {
		levellog.Controller("权限不够，无权访问用户列表")
		response.Failed(context, "权限不够，无权访问")
		return
	}

	// 权限校验通过之后获取参数
	var (
		p          simpletool.Page
		logoutUser models.LogoutUser
		wg         sync.WaitGroup
	)
	passChan := make(chan bool, 2)
	wg.Add(2)

	go func() {
		defer wg.Done()
		err = context.ShouldBindQuery(&p)
		if err != nil {
			w := fmt.Sprintf("参数page校验不通过, err: %s", err.Error())
			levellog.Controller(w)
			passChan <- false
			context.Abort()
			return
		}
		passChan <- true
	}()

	go func() {
		defer wg.Done()
		err = context.ShouldBindQuery(&logoutUser)
		if err != nil {
			levellog.Controller(fmt.Sprintf("获取注销用户参数失败, logoutUser: %v", logoutUser))
			passChan <- false
			context.Abort()
			return
		}
		passChan <- true
	}()

	wg.Wait()
	pass1 := <-passChan
	pass2 := <-passChan
	if !pass1 || !pass2 {
		w := "参数校验不通过，请检查参数"
		levellog.Controller(w)
		response.Failed(context, w)
		return
	}

	users, total, err := lgtusersvc.LogoutList(p, logoutUser)
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
