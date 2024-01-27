package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"server/common"
	"server/request"
	"server/response"
	"server/utils"
)

func (l LoginController) validation(req request.LoginRequest) bool {
	if req.Username == "" || req.Password == "" || utils.PromiseToInt(req.Promise) == 0 {
		return false
	}
	return true
}

func (l LoginController) Login(context *gin.Context) {
	var (
		loginRequest request.LoginRequest
		myErr        *common.MyError
	)
	if context.ShouldBind(&loginRequest) != nil {
		log.Println("绑定失败")
		response.Failed(context, "参数错误")
		return
	}
	ok := l.validation(loginRequest)
	if !ok {
		log.Println("参数校验不通过")
		response.Failed(context, "参数校验不通过")
		return
	}
	// 登录业务
	token, err := loginService.Login(loginRequest)
	if errors.As(err, &myErr) {
		log.Println("登录失败")
		response.FailedWithError(context, myErr)
		return
	}
	log.Println("token: ", token)
	response.Success(context, token, "登录成功")
}
