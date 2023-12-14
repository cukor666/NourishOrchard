package dao

import (
	"Gin/app"
	"Gin/models"
	"Gin/vo"
	"log"
)

type AdminDao struct{}

// SelectAllAdminByLimit 查询所有管理员，分页查询
func (ad AdminDao) SelectAllAdminByLimit(pageSize, currentPage int) (users []models.User, result vo.Result) {
	// 不把密码查出来给前端，防止密码泄露
	tx := app.MySQLDB().Omit("password", "DeletedAt").Where("promise = ?", 2)
	result.Rows = tx.Find(&[]models.User{}).RowsAffected
	result.Error = tx.Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&users).Error
	return
}

// AddAdmin 添加管理员
func (ad AdminDao) AddAdmin(ids []uint) error {
	return app.MySQLDB().Exec("UPDATE users SET `promise` = ? WHERE id IN ?", 2, ids).Error
}

// DeleteById 删除管理员
func (ad AdminDao) DeleteById(id uint) bool {
	err := app.MySQLDB().Table("users").Where("id = ?", id).Update("promise", 1).Error
	if err != nil {
		log.Printf("delete admin failed, err: %v\n", err)
		return false
	}
	return true
}
