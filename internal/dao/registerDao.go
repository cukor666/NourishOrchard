package dao

import (
	"fmt"
	"nourish-orchard-internal/models"
)

func (r RegisterDao) Register(account models.Account, admin models.Admin) (uint, bool) {
	tx := mysqlDB.Begin()
	err := tx.Create(&account).Error
	if err != nil {
		tx.Rollback()
		fmt.Println("插入account表失败")
		return 0, false
	}
	err = tx.Create(&admin).Error
	if err != nil {
		tx.Rollback()
		fmt.Println("插入admin表失败")
		return 0, false
	}
	tx.Commit()
	return admin.ID, true
}

func (r RegisterDao) RegisterEmployee(account models.Account, employee models.Employee) (uint, bool) {
	tx := mysqlDB.Begin()
	err := tx.Create(&account).Error
	if err != nil {
		tx.Rollback()
		fmt.Println("插入account表失败")
		return 0, false
	}
	err = tx.Create(&employee).Error
	if err != nil {
		tx.Rollback()
		fmt.Println("插入employee表失败")
		return 0, false
	}
	tx.Commit()
	return employee.ID, true
}
