package loginsvc

import (
	"server/dao/admindao"
	"server/dao/empdao"
	"server/dao/userdao"
)

// 普通用户登录
func userLogin(username string) (id uint, err error) {
	id, err = userdao.GetUId(username)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 员工登录
func employeeLogin(username string) (id uint, err error) {
	id, err = empdao.GetUId(username)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 管理员登录
func adminLogin(username string) (id uint, err error) {
	id, err = admindao.GetUId(username)
	if err != nil {
		return 0, err
	}
	return id, nil
}
