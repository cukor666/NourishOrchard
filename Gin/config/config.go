package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Config 总配置
type Config struct {
	SystemConfig `yaml:"systemConfig"`
	MySQLConfig  `yaml:"mysql"`
}

var conf Config

// GetConfig 获取总配置
func GetConfig() Config {
	return conf
}

// InitConfig 初始化总配置
func InitConfig() {
	// 读取yaml文件
	yamlFile, err := os.ReadFile("config/config.yaml") // 相对exe生成的路径
	if err != nil {
		log.Printf("读取总配置文件失败, err:%v", err)
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Printf("解析失败 yaml.Unmarshal err:%v", err)
		panic(err)
	}
	//log.Println(conf)
}

func init() {
	InitConfig() // 初始化总配置
}
