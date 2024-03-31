package fruitctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/config"
	"server/controller"
	"server/controller/args/header"
	"server/response"
	"server/service/fruitsvc"
)

// Detail 水果详情
func Detail(context *gin.Context) {
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

	type reqStruct struct {
		ID uint `form:"id" binding:"required,gt=0"`
	}

	var req reqStruct
	err = context.ShouldBindQuery(&req)
	if err != nil {
		w := fmt.Sprintf("参数校验失败，err: %s", err.Error())
		levellog.Controller(w)
		response.Failed(context, w)
		return
	}

	res, err := fruitsvc.Detail(int(req.ID))
	if err != nil {
		levellog.Controller("水果详情请求接口错误")
		response.Failed(context, "水果详情请求接口错误")
		return
	}

	response.Success(context, res, "查询成功")
}
