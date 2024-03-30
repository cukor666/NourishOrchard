package actctrl

import (
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/config"
	"server/controller"
	cm "server/controller/args/claims"
	"server/controller/args/header"
	"server/controller/impl"
	mc "server/models/code"
	"server/response"
	"server/utils/promisetool"
)

/**
AccountController是关于个人管理的接口控制管理，对应的大批管理由对应的Controller管理
例如：用户列表由UserController自行管理
*/

// GetAccount 获取账户信息
/**
header:
	Bearer Token
	username: CZKJ991706348185
body: 空
*/
func GetAccount(context *gin.Context) {
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
	// 查询数据库，将对应的账户个人信息返回给前端
	username := claims[cm.Username]
	promise := promisetool.ToInt(claims[cm.Promise].(string))
	var i controller.Acters
	switch promise {
	case mc.USER:
		i = &impl.UserImpl{}
	case mc.EMPLOYEE:
		i = &impl.EmployeeImpl{}
	case mc.ADMIN:
		i = &impl.AdminImpl{}
	default:
		levellog.Controller("权限不对")
		response.Failed(context, "权限不对")
		return
	}
	i.Info(context, username.(string))
}
