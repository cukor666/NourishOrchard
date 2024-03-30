package fruitctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/common/simpletool"
	"server/config"
	"server/controller"
	"server/controller/args/header"
	"server/models"
	"server/response"
	"server/service/fruitsvc"
)

// List 查询水果列表
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
	_, err = config.ParseAndVerifyJWT(token)
	if err != nil {
		levellog.Controller("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}

	var (
		p     simpletool.Page
		fruit models.Fruit
	)

	err1 := context.ShouldBindQuery(&p)
	err2 := context.ShouldBindQuery(&fruit)
	if err1 != nil || err2 != nil {
		levellog.Controller("绑定参数解构错误")
		response.Failed(context, "参数错误")
		return
	}

	fruits, total, err := fruitsvc.List(p, fruit)
	if err != nil {
		levellog.Controller(fmt.Sprintf("查询水果列表失败，fruits = %v", fruits))
		response.Failed(context, "查询水果列表失败")
		return
	}
	response.Success(context, gin.H{
		"fruits": fruits,
		"total":  total,
	}, "查询水果列表成功")
}
