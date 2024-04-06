package actsvc

import (
	"server/dao/actdao"
	"server/models"
)

// Get 获取账号信息
func Get(username string, promise int) (account models.Account, err error) {
	return actdao.Get(username, promise)
}
