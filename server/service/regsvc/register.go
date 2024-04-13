package regsvc

import (
	"errors"
	"server/common/levellog"
	"server/dao/actdao"
	"server/models"
	"server/utils/gentool"
	"server/utils/pwdtool"
)

// Register 用户注册业务
func Register(account models.Account, user models.User) (result models.Account, ok bool) {
	usernameChan := make(chan string)
	pwdChan := make(chan string)
	errChan := make(chan error)
	go func() {
		username := gentool.GenUsername()
		usernameChan <- username
	}()

	go func() {
		password, err := pwdtool.GetPwd(account.Password)
		if err != nil {
			levellog.Service("加密失败")
			errChan <- errors.New("加密失败")
			return
		}
		errChan <- nil
		pwdChan <- string(password)
	}()

	if err := <-errChan; err != nil {
		levellog.Service(err.Error())
		return
	}

	username := <-usernameChan
	user.Username, account.Username = username, username
	account.Password = <-pwdChan

	return actdao.Insert(account, user)
}
