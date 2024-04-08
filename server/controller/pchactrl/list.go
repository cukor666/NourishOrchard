package pchactrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/common/simpletool"
	"server/controller"
	"server/models"
	mc "server/models/code"
	"server/response"
	"server/service/pchasvc"
	"sync"
)

// List 查询采购列表接口
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
		levellog.Controller("权限不够拒绝请求")
		response.Failed(context, "权限不够拒绝请求")
	}
}

func listHandle(ctx *gin.Context) {
	var (
		p        simpletool.Page
		purchase models.Purchase
		wg       sync.WaitGroup
	)
	gCnt := 2
	wg.Add(gCnt)
	errChan := make(chan error, gCnt)

	go func() {
		defer wg.Done()
		err := ctx.ShouldBindQuery(&p)
		if err != nil {
			w := fmt.Sprintf("校验参数p失败, err: %s", err.Error())
			levellog.Controller(w)
			errChan <- err
			return
		}
		errChan <- nil
	}()

	go func() {
		defer wg.Done()
		err := ctx.ShouldBindQuery(&purchase)
		if err != nil {
			w := fmt.Sprintf("校验参数purchase失败, err: %s", err.Error())
			levellog.Controller(w)
			errChan <- err
			return
		}
		errChan <- nil
	}()

	wg.Wait()
	close(errChan)
	for err := range errChan {
		if err != nil {
			w := fmt.Sprintf("参数校验失败, err: %s", err.Error())
			levellog.Controller(w)
			response.Failed(ctx, w)
			return
		}
	}
	purchases, total, err := pchasvc.List(p, purchase)
	if err != nil {
		levellog.Controller("查询失败")
		response.Failed(ctx, "查询失败")
		return
	}
	response.Success(ctx, gin.H{
		"total":     total,
		"purchases": purchases,
	}, "查询成功")
}
