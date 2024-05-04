package fruitcmdyctrl

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/response"
	"server/service/fruitcmdysvc"
)

func List(context *gin.Context) {
	// 不需要做权限校验，因为每个人都可以查看列表
	// 直接返回所有水果数据即可
	var res FruitCmdyResponse
	var fruits []FruitCmdyResponse
	result, total, err := fruitcmdysvc.List()
	if err != nil {
		levellog.Controller("查询失败")
		response.Failed(context, "查询失败")
		return
	}

	for _, v := range result {
		err := res.set(v)
		if err != nil {
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
