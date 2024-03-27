package pwdmg

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// GetPassword 根据账号获取数据库中的密码
func GetPassword(username string) (password string, err error) {
	err = dao.GetMySQL().Model(&models.Account{}).Select("password").
		Where("username = ?", username).Take(&password).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("获取账户%s密码失败", username))
		return "", err
	}
	return password, nil
}
