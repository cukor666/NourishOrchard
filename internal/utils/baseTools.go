package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

const (
	prefix = "CZKJ"
)

// GetPwd 给密码加密
func GetPwd(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

// GenUsername 生成唯一账号名
func GenUsername() string {
	// 获取本地时间
	// now := time.Now()
	//timeStr := now.Format("2006-01-02 15:04:05")

	// 获取当前本地时间戳
	now := time.Now().Unix()
	fmt.Println("当前本地时间：", now)
	i := strconv.Itoa(rand.Int() % 100)
	timeStr := strconv.Itoa(int(now))
	return prefix + i + timeStr
}
