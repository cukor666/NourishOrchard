package empsvc

import (
	"errors"
	"regexp"
	"server/dao/empdao"
	"server/models"
)

// Info 根据账号获取员工信息
func Info(username string) (emp models.Employee, err error) {
	r := regexp.MustCompile("CZKJ[0-9]{3,}")
	matchString := r.MatchString(username)
	if !matchString {
		return models.Employee{}, errors.New("账号不符合系统规范")
	}
	return empdao.SelectByUsername(username)
}
