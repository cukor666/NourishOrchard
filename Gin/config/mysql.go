package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
)

type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

// GetDSN 获取mysql的dsn
func GetDSN() string {
	// 读取mysql.yaml文件
	yamlFile, err := os.ReadFile("config/mysql.yaml") // 相对exe生成的路径
	if err != nil {
		log.Printf("读取mysql.yaml文件失败, err:%v", err)
		panic(err)
	}
	var mysqlConfig MySQLConfig
	err = yaml.Unmarshal(yamlFile, &mysqlConfig)
	if err != nil {
		log.Printf("解析失败 yaml.Unmarshal err:%v", err)
		panic(err)
	}
	return mysqlConfig.User + ":" + mysqlConfig.Password +
		"@tcp(" + mysqlConfig.Host + ":" + strconv.Itoa(mysqlConfig.Port) + ")/" +
		mysqlConfig.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
}
