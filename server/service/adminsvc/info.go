package adminsvc

import (
	"errors"
	"regexp"
	"server/dao/admindao"
	"server/models"
)

// Info 根据账号获取管理员信息
func Info(username string) (admin models.Admin, err error) {
	r := regexp.MustCompile("CZKJ[0-9]{3,}")
	matchString := r.MatchString(username)
	if !matchString {
		return models.Admin{}, errors.New("账号不符合系统规范")
	}
	return admindao.SelectByUsername(username)
}
