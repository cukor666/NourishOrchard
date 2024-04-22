package userdao

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// SelectByUsername 通过username查询查询用户
func SelectByUsername(username string) (user models.User, err error) {
	db := dao.GetMySQL()
	err = db.Model(&user).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// SelectByUsernameAndPhone 通过账号和电话查找用户信息
func SelectByUsernameAndPhone(username, phone string) (user models.User, err error) {
	db := dao.GetMySQL()
	err = db.Model(&models.User{}).Where("username = ? AND phone = ?", username, phone).Take(&user).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("查询用户失败, user = %v", user))
		return models.User{}, err
	}
	return user, nil
}

// ExistsById 通过id查询用户是否存在
func ExistsById(id int64) bool {
	db := dao.GetMySQL()
	var count int64
	err := db.Model(&models.User{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false
	}
	if count > 0 {
		return true
	}
	return false
}
