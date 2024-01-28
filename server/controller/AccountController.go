package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/config"
	"server/response"
	"server/utils"
)

func (a AccountController) GetAccount(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader("Authorization")
	token, _ := utils.GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	// 解析token，或token里面的内容
	jwt, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		log.Println("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	log.Println(jwt)
	response.Success(context, jwt, "解析JWT成功")
}
