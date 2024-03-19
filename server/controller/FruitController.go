package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/config"
	"server/controller/args/header"
	"server/response"
	"server/service"
)

// List 查询水果列表
func (f FruitController) List(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levelLog("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	_, err = config.ParseAndVerifyJWT(token)
	if err != nil {
		levelLog("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	fruits, err := service.FruitService{}.List()
	if err != nil {
		levelLog(fmt.Sprintf("查询水果列表失败，fruits = %v", fruits))
		response.Failed(context, "查询水果列表失败")
		return
	}
	response.Success(context, fruits, "查询水果列表成功")
}
