package utils

import (
	"golang.org/x/crypto/bcrypt"
	mc "server/models/code"
	"strings"
)

func PromiseToString(promise int) string {
	switch promise {
	case mc.USER:
		return "user"
	case mc.EMPLOYEE:
		return "employee"
	case mc.ADMIN:
		return "admin"
	}
	return ""
}

func PromiseToInt(promise string) int {
	promise = strings.ToLower(promise)
	switch promise {
	case "user":
		return mc.USER
	case "employee":
		return mc.EMPLOYEE
	case "admin":
		return mc.ADMIN
	}
	return 0
}

// GetPwd 给密码加密
func GetPwd(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

// PwdOK 校验密码是否正确
func PwdOK(origin, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(origin), []byte(password)) == nil
}
