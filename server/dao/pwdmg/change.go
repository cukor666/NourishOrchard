package pwdmg

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// ChangePassword 修改密码
func ChangePassword(username, password string) (err error) {
	err = dao.GetMySQL().Model(&models.Account{}).Where("username = ?", username).Update("password", password).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("用户%s更新密码失败", username))
		return err
	}
	return nil
}
