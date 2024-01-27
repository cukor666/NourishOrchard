package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"server/config"
)

var mysqlDB *gorm.DB

type AccountDao struct{}
type UserDao struct{}

// 初始化数据源，以连接数据库
func init() {
	dsn := config.GetDSN()
	var mysqlLogger = logger.Default.LogMode(logger.Info)
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger, // 添加SQL日志
	})
	if err != nil {
		panic(fmt.Sprintf("gorm.Open failed, myerr: %v", err))
	}
	mysqlDB = open
}
