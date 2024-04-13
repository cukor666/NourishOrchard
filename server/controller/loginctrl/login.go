package loginctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/response"
	"server/service/loginsvc"
)

// Login 登录接口
func Login(context *gin.Context) {
	var req loginRequest // 定义登录请求结构
	if err := context.ShouldBind(&req); err != nil {
		w := fmt.Sprintf("参数校验失败: %v", err)
		levellog.Controller(w)
		response.Failed(context, w)
		return
	}

	// 登录业务
	token, err := loginsvc.Login(req.toAccount())
	if err != nil {
		w := fmt.Sprintf("登录失败: %v", err)
		levellog.Controller(w)
		response.Failed(context, w)
		return
	}
	levellog.Controller(fmt.Sprintf("token: %v", token))
	response.Success(context, token, "登录成功")
}
