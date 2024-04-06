package actdao

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// Insert
// 添加账户
func Insert(account models.Account, user models.User) (uint, bool) {
	db := dao.GetMySQL()
	// 开启事务
	tx := db.Begin()
	// 将数据插入到account表中
	err := tx.Create(&account).Error
	if err != nil {
		levellog.Dao("数据添加到account表失败")
		tx.Rollback()
		return 0, false
	}
	// 将数据插入到user表中
	err = tx.Create(&user).Error
	if err != nil {
		levellog.Dao("数据添加到user表失败")
		tx.Rollback()
		return 0, false
	}
	if err = tx.Commit().Error; err != nil {
		levellog.Dao(fmt.Sprintf("err: %v", err))
		return 0, false
	}
	return user.ID, true
}
