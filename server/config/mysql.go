package config

import (
	"fmt"
)

// MySQLConfig MySQL配置
type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Param    string `yaml:"param"`
}

// GetDSN 获取mysql的dsn
func GetDSN() string {
	m := conf.MySQLConfig
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.User, m.Password, m.Host, m.Port, m.DBName, m.Param)
}
