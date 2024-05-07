package fruitcmdyctrl

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/response"
	"server/service/fruitcmdysvc"
	"strconv"
)

func List(context *gin.Context) {
	// 不需要做权限校验，因为每个人都可以查看列表
	// 直接返回所有水果数据即可
	var res FruitCmdyResponse
	var fruits []FruitCmdyResponse
	// 标识南北与全部水果 0:全部 1:南方 2:北方
	value := context.Query("flag")
	flag, err := strconv.Atoi(value)
	if err != nil || flag < 0 || flag > 2 {
		levellog.Controller("参数错误")
		response.Failed(context, "参数错误")
		return
	}
	result, total, err := fruitcmdysvc.List(flag)
	if err != nil {
		levellog.Controller("查询失败")
		response.Failed(context, "查询失败")
		return
	}

	for _, v := range result {
		if err := res.set(v); err != nil {
			levellog.Controller("数据转换失败")
			response.Failed(context, "数据转换失败")
			return
		}
		fruits = append(fruits, res)
	}

	response.Success(context, gin.H{
		"fruits": fruits,
		"total":  total,
	}, "查询成功")
}
