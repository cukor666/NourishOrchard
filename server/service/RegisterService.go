package service

import (
	"log"
	"server/models"
	"server/request"
	"server/response"
	"server/utils"
)

func (r RegisterService) deconstruct(req request.RegisterRequest) (models.Account, models.User) {
	return models.Account{
			Password: req.Password,
			Promise:  req.Promise,
		}, models.User{
			Name:     req.Name,
			Gender:   req.Gender,
			Phone:    req.Phone,
			Address:  req.Address,
			Birthday: req.Birthday,
		}
}

func (r RegisterService) Register(req request.RegisterRequest) (response.RegisterResponse, bool) {
	account, user := r.deconstruct(req)
	username := utils.GenUsername()
	user.Username = username
	uid, userOk := userDao.Insert(user)
	if !userOk {
		return response.RegisterResponse{}, false
	}
	account.Username = username
	password, err := utils.GetPwd(account.Password)
	if err != nil {
		log.Println("加密失败")
		return response.RegisterResponse{}, false
	}
	account.Password = string(password)
	accountOk := accountDao.Insert(account)
	if !accountOk {
		return response.RegisterResponse{}, false
	}
	return response.RegisterResponse{
		ID:       uid,
		Username: account.Username,
		Promise:  utils.PromiseToString(account.Promise),
	}, true
}
