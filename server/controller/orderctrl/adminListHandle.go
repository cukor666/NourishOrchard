package orderctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/common/simpletool"
	"server/models"
	"server/response"
	"server/service/ordersvc"
	"sync"
)

func listHandle(ctx *gin.Context) {
	var (
		p     simpletool.Page
		order models.Order
		wg    sync.WaitGroup
	)
	wg.Add(2)
	errChan := make(chan error, 2)

	go func() {
		defer wg.Done()
		if err := ctx.ShouldBindQuery(&p); err != nil {
			w := fmt.Sprintf("参数校验失败，err: %s", err.Error())
			levellog.Controller(w)
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := ctx.ShouldBindQuery(&order); err != nil {
			w := fmt.Sprintf("参数校验失败，err: %s", err.Error())
			levellog.Controller(w)
			errChan <- err
			return
		}
	}()

	wg.Wait()
	close(errChan)
	for err := range errChan {
		if err != nil {
			response.Failed(ctx, "参数校验失败")
			return
		}
	}
	orders, total, err := ordersvc.List(p, order)
	if err != nil {
		levellog.Controller(err.Error())
		response.Failed(ctx, "获取订单列表失败")
		return
	}
	response.Success(ctx, gin.H{
		"total":  total,
		"orders": orders,
	}, "查询成功")
}
