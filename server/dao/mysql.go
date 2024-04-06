package dao

import "gorm.io/gorm"

var mysqlDB *gorm.DB

func GetMySQL() *gorm.DB {
	return mysqlDB
}
