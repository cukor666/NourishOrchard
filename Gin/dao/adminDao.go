package dao

import (
	"Gin/moudels"
	"Gin/utils"
)

type AdminDao struct{}

// 查询所有管理员，分页查询
func (ad AdminDao) SelectAllAdminByLimt(pageSize, currentPage int) (users []moudels.User, result utils.Result) {
	// 不把密码查出来给前端，防止密码泄露
	tx := db.Omit("password", "DeletedAt").Where("promise = ?", 2)
	result.Rows = tx.Find(&[]moudels.User{}).RowsAffected
	result.Error = tx.Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&users).Error
	return
}

// 添加管理员
func (ad AdminDao) AddAdmin(ids []uint) error {
	// 官方示例： db.Exec("UPDATE orders SET shipped_at = ? WHERE id IN ?", time.Now(), []int64{1, 2, 3})
	err := db.Exec("UPDATE users SET `promise` = ? WHERE id IN ?", 2, ids).Error
	return err
}
