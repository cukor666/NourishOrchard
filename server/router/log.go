package router

import (
	"log"
	"os"
	"time"
)

func newLogFile() {
	// 创建一个新的日志文件
	file, err := os.OpenFile("log/app.log."+time.Now().Format("2006-01-02"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// 将文件设置为 log 包的输出目标
	log.SetOutput(file)
}
