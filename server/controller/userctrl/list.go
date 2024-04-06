package userctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/common/simpletool"
	"server/controller"
	"server/controller/spliterr"
	"server/models"
	mc "server/models/code"
	"server/response"
	"server/service/usersvc"
	"sync"
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
func List(context *gin.Context) {
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限校验失败")
		response.Failed(context, "权限校验失败")
		return
	}
	// 获取请求方的权限，如果是普通用户权限则无权访问
	promise := controller.GetPromise(claims)
	if promise < mc.EMPLOYEE {
		levellog.Controller("权限不够，无权访问用户列表")
		response.Failed(context, "权限不够，无权访问")
		return
	}

	// 权限校验通过之后获取参数
	var (
		p    simpletool.Page
		user models.User
		wg   sync.WaitGroup
	)

	wg.Add(2)
	passChan := make(chan bool, 2)

	// 获取分页参数
	go func() {
		defer wg.Done()
		err = context.ShouldBindQuery(&p)
		if err != nil {
			msg := spliterr.GetErrMsg(err.Error())
			w := fmt.Sprintf("分页参数校验失败, err: %s", msg)
			levellog.Controller(w)
			passChan <- false
			context.Abort()
			return
		}
		passChan <- true
	}()

	// 获取查询用户的参数
	go func() {
		defer wg.Done()
		err = context.ShouldBindQuery(&user)
		if err != nil {
			msg := spliterr.GetErrMsg(err.Error())
			w := fmt.Sprintf("获取用户参数失败, err: %s", msg)
			levellog.Controller(w)
			passChan <- false
			context.Abort()
			return
		}
		levellog.Controller(fmt.Sprintf("user = %v", user))
		passChan <- true
	}()

	wg.Wait()
	g1 := <-passChan
	g2 := <-passChan
	if !g1 || !g2 {
		levellog.Controller("参数校验失败")
		response.Failed(context, "参数错误")
		return
	}

	users, total, err := usersvc.List(p, user)
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
