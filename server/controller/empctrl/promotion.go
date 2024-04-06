package empctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/config"
	"server/controller"
	cm "server/controller/args/claims"
	"server/controller/args/header"
	"server/controller/spliterr"
	"server/request"
	"server/response"
	"server/service/actsvc"
	"server/service/empsvc"
	"server/utils/promisetool"
	"server/utils/pwdtool"
)

// Promotion 晋升管理员
func Promotion(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := controller.GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levellog.Controller("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		levellog.Controller("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	reqUsername := claims[cm.Username].(string)
	reqPromise := promisetool.ToInt(claims[cm.Promise].(string))
	account, err := actsvc.Get(reqUsername, reqPromise)
	if err != nil {
		levellog.Controller("获取账号信息失败")
		response.Failed(context, "获取账号信息失败")
		return
	}

	var req request.PromotionRequest
	err = context.ShouldBindJSON(&req)
	if err != nil {
		levellog.Controller(fmt.Sprintf("数据绑定失败, promotionReq = %v", req))
		msg := spliterr.GetErrMsg(err.Error())
		levellog.Controller(fmt.Sprintf("err: %s", msg))
		response.Failed(context, msg)
		return
	}

	if !pwdtool.PwdOK(account.Password, req.Password) {
		levellog.Controller(fmt.Sprintf("请求账号%s密码错误，拒绝请求", reqUsername))
		response.Failed(context, "密码错误")
		return
	}

	err = empsvc.Promotion(req)
	if err != nil {
		levellog.Controller("晋升管理员失败")
		response.Failed(context, "晋升管理员失败")
		return
	}
	response.Success(context, 0, "晋升管理员成功")
}
