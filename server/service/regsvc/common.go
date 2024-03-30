package regsvc

import (
	"server/models"
	mc "server/models/code"
	"server/request"
)

/*
*
解构，对前端参数解构
返回User对象
*/
func deconstruct(req request.RegisterRequest) (models.Account, models.User) {
	return models.Account{
			Password: req.Password,
			//Promise:  req.Promise,
			Promise: mc.USER, // 该接口只提供用户注册，管理员和员工的注册由内部自己完成
		}, models.User{
			Name:     req.Name,
			Gender:   req.Gender,
			Phone:    req.Phone,
			Address:  req.Address,
			Birthday: req.Birthday,
		}
}
