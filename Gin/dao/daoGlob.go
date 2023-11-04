package dao

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 全局数据库对象
var db *gorm.DB

// 初始化数据源，以连接数据库
func init() {
	dsn := "root:cukor@tcp(localhost:3306)/nourish_orchard?charset=utf8mb4&parseTime=True&loc=Local"
	var mysqlLogger = logger.Default.LogMode(logger.Info)
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger, // 添加SQL日志
	})
	if err != nil {
		log.Printf("gorm.Open failed, err: %v\n", err)
		return
	}
	db = open
}
