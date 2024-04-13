package loginctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/controller/spliterr"
	"server/response"
	"server/service/loginsvc"
)

// Login 登录接口
func Login(context *gin.Context) {
	var req loginRequest // 定义登录请求结构
	if err := context.ShouldBind(&req); err != nil {
		levellog.Controller(spliterr.GetErrMsg(err.Error()))
		response.Failed(context, "参数校验失败")
		return
	}

	// 登录业务
	token, err := loginsvc.Login(req.toAccount())
	if err != nil {
		levellog.Controller("登录失败")
		response.Failed(context, "登录失败")
		return
	}
	levellog.Controller(fmt.Sprintf("token: %v", token))
	response.Success(context, token, "登录成功")
}
