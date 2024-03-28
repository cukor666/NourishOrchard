package service

import (
	"server/common/levellog"
	"server/models"
	mc "server/models/code"
	"server/request"
	"server/response"
	"server/utils/gentool"
	"server/utils/promisetool"
	"server/utils/pwdtool"
)

/*
*
解构，对前端参数解构
返回User对象
*/
func (r RegisterService) deconstruct(req request.RegisterRequest) (models.Account, models.User) {
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

// Register 用户注册业务
func (r RegisterService) Register(req request.RegisterRequest) (response.RegisterResponse, bool) {
	account, user := r.deconstruct(req)
	username := gentool.GenUsername()
	user.Username = username
	account.Username = username
	password, err := pwdtool.GetPwd(account.Password)
	if err != nil {
		levellog.Service("加密失败")
		return response.RegisterResponse{}, false
	}
	account.Password = string(password)
	uid, ok := accountDao.Insert(account, user)
	if !ok {
		return response.RegisterResponse{}, false
	}
	return response.RegisterResponse{
		ID:       uid,
		Username: account.Username,
		Promise:  promisetool.ToString(account.Promise),
	}, true
}
