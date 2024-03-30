package regsvc

import (
	"server/common/levellog"
	"server/dao/actdao"
	"server/request"
	"server/response"
	"server/utils/gentool"
	"server/utils/promisetool"
	"server/utils/pwdtool"
)

// Register 用户注册业务
func Register(req request.RegisterRequest) (response.RegisterResponse, bool) {
	account, user := deconstruct(req)
	username := gentool.GenUsername()
	user.Username = username
	account.Username = username
	password, err := pwdtool.GetPwd(account.Password)
	if err != nil {
		levellog.Service("加密失败")
		return response.RegisterResponse{}, false
	}
	account.Password = string(password)
	uid, ok := actdao.Insert(account, user)
	if !ok {
		return response.RegisterResponse{}, false
	}
	return response.RegisterResponse{
		ID:       uid,
		Username: account.Username,
		Promise:  promisetool.ToString(account.Promise),
	}, true
}
