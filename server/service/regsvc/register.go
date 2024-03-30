package regsvc

import (
	"errors"
	"server/common/levellog"
	"server/dao/actdao"
	"server/request"
	"server/response"
	"server/utils/gentool"
	"server/utils/promisetool"
	"server/utils/pwdtool"
	"sync"
)

// Register 用户注册业务
func Register(req request.RegisterRequest) (response.RegisterResponse, bool) {
	account, user := deconstruct(req)
	var wg sync.WaitGroup
	wg.Add(2)
	usernameChan := make(chan string)
	pwdChan := make(chan string)
	errChan := make(chan error)
	go func() {
		defer wg.Done()
		username := gentool.GenUsername()
		usernameChan <- username
	}()

	go func() {
		defer wg.Done()
		password, err := pwdtool.GetPwd(account.Password)
		if err != nil {
			levellog.Service("加密失败")
			errChan <- errors.New("加密失败")
		}
		errChan <- nil
		pwdChan <- string(password)
	}()

	err := <-errChan
	if err != nil {
		levellog.Service(err.Error())
		return response.RegisterResponse{}, false
	}

	username := <-usernameChan
	user.Username, account.Username = username, username
	account.Password = <-pwdChan

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
