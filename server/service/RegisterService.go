package service

import (
	"log"
	"server/common"
	"server/models"
	"server/request"
	"server/response"
	"server/utils"
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
			Promise: common.USER, // 该接口只提供用户注册，管理员和员工的注册由内部自己完成
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
	username := utils.GenUsername()
	user.Username = username
	account.Username = username
	password, err := utils.GetPwd(account.Password)
	if err != nil {
		log.Println("加密失败")
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
		Promise:  utils.PromiseToString(account.Promise),
	}, true
}
