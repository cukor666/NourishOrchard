package actsvc

import (
	"regexp"
	"server/dao/actdao"
	"server/utils/gentool"
)

// IsExists 判断账号在数据库中是否存在
func IsExists(username string, promise int) bool {
	// 如果格式不对的话直接返回false
	prefix := gentool.GetUsernamePrefix()
	compile := regexp.MustCompile("^" + prefix + "[0-9]+$")
	ok := compile.MatchString(username)
	if !ok {
		return false
	}
	// 开始查询数据库
	cnt := actdao.GetCountByUsername(username, promise)
	return cnt != 0
}
