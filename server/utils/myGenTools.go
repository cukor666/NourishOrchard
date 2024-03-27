package utils

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	prefix = "CZKJ"
)

func GetUsernamePrefix() string {
	return prefix
}

// GenUsername 生成唯一账号名
func GenUsername() string {
	// 获取当前本地时间戳
	now := time.Now().Unix()
	i := strconv.Itoa(rand.Int() % 100)
	timeStr := strconv.Itoa(int(now))
	return prefix + i + timeStr
}
