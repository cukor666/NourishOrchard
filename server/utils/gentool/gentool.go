package gentool

import (
	"math/rand"
	"strconv"
	"time"
)

// GenUsername 生成唯一账号名
func GenUsername() string {
	// 获取当前本地时间戳
	now := time.Now().Unix()
	i := strconv.Itoa(rand.Int() % 100)
	timeStr := strconv.Itoa(int(now))
	return "CZKJ" + i + timeStr
}
