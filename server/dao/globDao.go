package dao

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"server/config"
)

var mysqlDB *gorm.DB
var redisDB = config.GetConfig().RedisConfig.GetClient()

type AccountDao struct{}
type UserDao struct{}
type AdminDao struct{}
type EmployeeDao struct{}
type FruitDao struct{}

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

// GetRedisDB 开放给外部访问redis实例
func GetRedisDB() *redis.Client {
	return redisDB
}
