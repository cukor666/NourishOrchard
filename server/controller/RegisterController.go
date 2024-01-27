package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/request"
	"server/response"
	"server/utils"
)

// 数据校验
func (r RegisterController) validation(rg request.RegisterRequest) bool {
	if rg.Name == "" || rg.Password == "" || utils.PromiseToString(rg.Promise) == "myerr" {
		return false
	}
	return true
}

// Register 注册
func (r RegisterController) Register(context *gin.Context) {
	var (
		registerRequest  request.RegisterRequest
		registerResponse response.RegisterResponse
	)
	if context.ShouldBind(&registerRequest) != nil {
		log.Println("绑定参数失败")
		response.Failed(context, "绑定参数失败")
		return
	}
	log.Println("绑定成功")
	log.Println(registerRequest)

	ok := r.validation(registerRequest)
	if !ok {
		response.Failed(context, "参数错误")
		return
	}
	// 数据校验通过，交给业务层
	registerResponse, ok = registerService.Register(registerRequest)
	if !ok {
		log.Println("注册失败")
		response.FailedWithCode(context, 500, "注册失败")
		return
	}
	response.Success(context, registerResponse, "绑定成功")
}
