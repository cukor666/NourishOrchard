package supctrl

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
	"server/service/supsvc"
	"server/utils/promisetool"
	"sync"
)

func List(context *gin.Context) {
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
	if promise < mc.EMPLOYEE {
		levellog.Controller("权限不够，无权访问用户列表")
		response.Failed(context, "权限不够，无权访问")
		return
	}

	// 权限校验通过之后获取参数
	var (
		p    simpletool.Page
		sups models.Supplier
		wg   sync.WaitGroup
	)
	wg.Add(2)
	passChan := make(chan bool, 2)

	go func() {
		defer wg.Done()
		err := context.ShouldBindQuery(&p)
		if err != nil {
			w := fmt.Sprintf("参数page校验失败，err: %s", err.Error())
			levellog.Controller(w)
			passChan <- false
			context.Abort()
			return
		}
		passChan <- true
	}()

	go func() {
		defer wg.Done()
		err := context.ShouldBindQuery(&sups)
		if err != nil {
			w := fmt.Sprintf("参数supplier校验失败，err: %s", err.Error())
			levellog.Controller(w)
			passChan <- false
			context.Abort()
		}
		passChan <- true
	}()

	wg.Wait()
	ps1 := <-passChan
	ps2 := <-passChan
	if !ps1 || !ps2 {
		levellog.Controller("参数校验不通过")
		response.Failed(context, "参数错误")
		return
	}

	suppliers, total, err := supsvc.List(p, sups)
	if err != nil {
		levellog.Controller("参数错误，该数据不存在")
		response.Failed(context, "参数错误，该数据不存在，请检查")
		return
	}
	response.Success(context, gin.H{
		"suppliers": suppliers,
		"total":     total,
	}, "查询供应商列表成功")
}
