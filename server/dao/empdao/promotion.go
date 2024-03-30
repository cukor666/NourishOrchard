package empdao

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
	mc "server/models/code"
)

// Promotion 晋升管理员
func Promotion(username, name, email, mark string) error {
	db := dao.GetMySQL()
	tx := db.Begin()
	a := models.Admin{
		Username: username,
		Name:     name,
		Email:    email,
	}

	// 更新account表中的对应账号的权限
	err := tx.Model(&models.Account{}).Where("username = ? AND promise = ?", username, mc.EMPLOYEE).
		Update("promise", mc.ADMIN).Error

	if err != nil {
		var am models.Account
		str := fmt.Sprintf("从%s表中更新username = %s的权限失败", am.TableName(), username)
		levellog.Dao(str)
		tx.Rollback()
		return err
	}

	err = tx.Model(&models.Admin{}).Create(&a).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("将%v插入到%s表中失败", a, a.TableName()))
		tx.Rollback()
		return err
	}

	// 修改employee_status表的状态
	err = tx.Model(&models.EmployeeStatus{}).Where("username = ?", username).Updates(map[string]interface{}{
		"status": mc.TransferFromAPosition,
		"mark":   mark,
	}).Error

	if err != nil {
		var esm models.EmployeeStatus
		levellog.Dao(fmt.Sprintf("在%s表中修改username = %s失败", esm.TableName(), username))
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
