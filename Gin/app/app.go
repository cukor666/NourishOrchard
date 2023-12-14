package app

import "gorm.io/gorm"

var (
	mysqlDB *gorm.DB // MySQL连接
)

// MySQLDB 获取MySQL数据库连接
func MySQLDB() *gorm.DB {
	return mysqlDB
}

// SetMySQLDB 设置MySQL数据库连接
func SetMySQLDB(db *gorm.DB) bool {
	mysqlDB = db
	return mysqlDB != nil
}
