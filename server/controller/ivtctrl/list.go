package ivtctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/common/simpletool"
	"server/controller"
	"server/models"
	mc "server/models/code"
	"server/response"
	"server/service/ivtsvc"
	"sync"
)

// List 查询库存信息接口
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
		levellog.Controller("权限不够，拒绝请求")
		response.Failed(context, "权限不够拒绝请求")
	}
}

func listHandle(ctx *gin.Context) {
	var (
		p   simpletool.Page
		ivt models.Inventory
		wg  sync.WaitGroup
	)
	wg.Add(2)
	errChan := make(chan error, 2)

	go func() {
		defer wg.Done()
		err := ctx.ShouldBindQuery(&p)
		if err != nil {
			w := fmt.Sprintf("参数page校验不通过, err: %s", err.Error())
			levellog.Controller(w)
			response.Failed(ctx, "参数page传递错误")
			errChan <- err
			return
		}
		errChan <- nil
	}()

	go func() {
		defer wg.Done()
		err := ctx.ShouldBindQuery(&ivt)
		if err != nil {
			w := fmt.Sprintf("参数ivt校验不通过, err: %s", err.Error())
			levellog.Controller(w)
			response.Failed(ctx, "参数ivt传递错误")
			errChan <- err
			return
		}
		errChan <- nil
	}()

	wg.Wait()
	err1 := <-errChan
	err2 := <-errChan
	if err1 != nil || err2 != nil {
		levellog.Controller("参数传递错误")
		response.Failed(ctx, "参数传递错误")
		return
	}

	inventors, total, err := ivtsvc.List(p, ivt)
	if err != nil {
		levellog.Controller("查询失败")
		response.Failed(ctx, "查询失败")
		return
	}
	response.Success(ctx, gin.H{
		"total":     total,
		"inventors": inventors,
	}, "查询成功")
}
