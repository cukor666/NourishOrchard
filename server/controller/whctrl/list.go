package whctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/common/simpletool"
	"server/controller"
	"server/models"
	mc "server/models/code"
	"server/response"
	"server/service/whsvc"
	"sync"
)

// List 查询仓库列表
func List(context *gin.Context) {
	claims, err := controller.ValidAuth(context)
	if err != nil {
		levellog.Controller("权限校验失败")
		response.Failed(context, "权限校验失败")
		return
	}
	promise := controller.GetPromise(claims)
	switch promise {
	case mc.ADMIN:
		listHandle(context)
	default:
		levellog.Controller("权限不够")
		response.Failed(context, "权限不够，拒绝请求")
	}
}

func listHandle(ctx *gin.Context) {
	var (
		p  simpletool.Page
		w  models.Warehouse
		wg sync.WaitGroup
	)
	wg.Add(2)
	errChan := make(chan error, 2)

	go func() {
		defer wg.Done()
		if err := ctx.ShouldBindQuery(&p); err != nil {
			levellog.Controller("参数page校验不通过")
			errChan <- err
			return
		}
		errChan <- nil
	}()

	go func() {
		defer wg.Done()
		if err := ctx.ShouldBindQuery(&w); err != nil {
			levellog.Controller("参数warehouse校验不通过")
			errChan <- err
			return
		}
		levellog.Controller(fmt.Sprintf("warehouse = %v", w))
		errChan <- nil
	}()

	wg.Wait()

	err1 := <-errChan
	err2 := <-errChan
	close(errChan)
	if err1 != nil || err2 != nil {
		levellog.Controller("参数不正确")
		response.Failed(ctx, "参数校验不通过拒绝请求，请检查参数")
		return
	}

	warehouses, total, err := whsvc.List(p, w)
	if err != nil {
		str := fmt.Sprintf("数据库中无该数据，page = %v, warehouse = %v", p, w)
		levellog.Controller(str)
		response.Failed(ctx, str)
		return
	}
	response.Success(ctx, gin.H{
		"total":      total,
		"warehouses": warehouses,
	}, "查询成功")
}
