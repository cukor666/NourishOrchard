package service

import (
	"regexp"
	"server/utils"
)

// IsExists 判断账号在数据库中是否存在
func (a AccountService) IsExists(username string, promise int) bool {
	// 如果格式不对的话直接返回false
	prefix := utils.GetUsernamePrefix()
	compile := regexp.MustCompile("^" + prefix + "[0-9]+$")
	ok := compile.MatchString(username)
	if !ok {
		return false
	}
	// 开始查询数据库
	cnt := accountDao.GetCountByUsername(username, promise)
	return cnt != 0
}
