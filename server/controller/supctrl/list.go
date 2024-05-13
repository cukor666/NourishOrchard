package supctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/common/simpletool"
	"server/controller"
	"server/models"
	mc "server/models/code"
	"server/response"
	"server/service/supsvc"
	"sync"
)

// List 供应商列表
func List(context *gin.Context) {
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限校验失败")
		response.Failed(context, "权限校验失败")
		return
	}

	promise := controller.GetPromise(claims)
	switch promise {
	case mc.EMPLOYEE, mc.ADMIN:
		listHandle(context)
	default:
		levellog.Controller("权限不够")
		response.Failed(context, "权限不够")
	}
}

func listHandle(ctx *gin.Context) {
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
		err := ctx.ShouldBindQuery(&p)
		if err != nil {
			w := fmt.Sprintf("参数page校验失败，err: %s", err.Error())
			levellog.Controller(w)
			passChan <- false
			ctx.Abort()
			return
		}
		passChan <- true
	}()

	go func() {
		defer wg.Done()
		err := ctx.ShouldBindQuery(&sups)
		if err != nil {
			w := fmt.Sprintf("参数supplier校验失败，err: %s", err.Error())
			levellog.Controller(w)
			passChan <- false
			ctx.Abort()
		}
		passChan <- true
	}()

	wg.Wait()
	ps1 := <-passChan
	ps2 := <-passChan
	if !ps1 || !ps2 {
		levellog.Controller("参数校验不通过")
		response.Failed(ctx, "参数错误")
		return
	}

	suppliers, total, err := supsvc.List(p, sups)
	if err != nil {
		levellog.Controller("参数错误，该数据不存在")
		response.Failed(ctx, "参数错误，该数据不存在，请检查")
		return
	}
	response.Success(ctx, gin.H{
		"suppliers": suppliers,
		"total":     total,
	}, "查询供应商列表成功")
}
