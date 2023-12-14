package dao

import (
	"Gin/app"
	"Gin/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 全局数据库对象
//var db *gorm.DB

// 初始化数据源，以连接数据库
func init() {
	//config.InitConfig()
	dsn := config.GetDSN()
	//log.Printf("dsn: %s", dsn)
	var mysqlLogger = logger.Default.LogMode(logger.Info)
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger, // 添加SQL日志
	})
	if err != nil {
		panic(fmt.Sprintf("gorm.Open failed, err: %v", err))
	}
	app.SetMySQLDB(open)
}
