package actdao

import (
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// Insert 添加账户
func Insert(account models.Account, user models.User) (r models.Account, ok bool) {
	// 开启事务
	tx := dao.GetMySQL().Begin()
	// 将数据插入到account表中
	if err := tx.Create(&account).Error; err != nil {
		levellog.Dao("数据添加到account表失败")
		tx.Rollback()
		return
	}
	// 将数据插入到user表中
	if err := tx.Create(&user).Error; err != nil {
		levellog.Dao("数据添加到user表失败")
		tx.Rollback()
		return
	}
	tx.Commit()

	r.Username, r.Promise = account.Username, account.Promise
	ok = true
	return
}
