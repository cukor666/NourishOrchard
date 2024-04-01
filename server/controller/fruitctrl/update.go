package fruitctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/config"
	"server/controller"
	cm "server/controller/args/claims"
	"server/controller/args/header"
	mc "server/models/code"
	"server/response"
	"server/service/fruitsvc"
	"server/utils/promisetool"
	"sync"
)

// Update 更新水果接口
func Update(context *gin.Context) {
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
	promise := promisetool.ToInt(claims[cm.Promise].(string))
	var fruit FruitUpdateReq
	switch promise {
	case mc.USER:
		levellog.Controller("权限不够")
		response.Failed(context, "权限不够")
	case mc.EMPLOYEE, mc.ADMIN:
		err = context.ShouldBindJSON(&fruit)
		if err != nil {
			w := fmt.Sprintf("数据绑定失败, err: %s", err.Error())
			levellog.Controller(w)
			response.Failed(context, w)
			return
		}
		// 自定义参数校验
		ok := validData(fruit)
		if !ok {
			levellog.Controller("参数校验失败，请检查参数")
			response.Failed(context, "参数校验失败请检查参数")
			return
		}

		err = fruitsvc.Update(fruit.ToFruit())
		if err != nil {
			levellog.Controller("更新失败")
			response.Failed(context, "更新失败")
			return
		}
		response.Success(context, 0, "更新成功")
	default:
		levellog.Controller("未知权限")
		response.Failed(context, "未知权限")
	}
}

func validData(req FruitUpdateReq) bool {
	defer func() {
		if err := recover(); err != nil {
			levellog.Controller(fmt.Sprintf("自定义校验时出错，可能是goroutine的问题, err: %v", err))
		}
	}()
	const goNum = 6
	oks := [goNum]bool{}
	var wg sync.WaitGroup
	wg.Add(goNum)
	go func() {
		defer wg.Done()
		oks[0] = validName(req.Name)
	}()

	go func() {
		defer wg.Done()
		oks[1] = validWater(req.Water)
	}()

	go func() {
		defer wg.Done()
		oks[2] = validSugar(req.Sugar)
	}()

	go func() {
		defer wg.Done()
		oks[3] = validShelfLife(req.ShelfLife)
	}()

	go func() {
		defer wg.Done()
		oks[4] = validOrigin(req.Origin)
	}()

	go func() {
		defer wg.Done()
		oks[5] = validSupplierId(req.SupplierId)
	}()

	wg.Wait()

	for _, ok := range oks {
		if !ok {
			return false
		}
	}
	return true
}
