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
