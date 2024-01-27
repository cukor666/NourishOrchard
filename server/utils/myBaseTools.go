package utils

import (
	"golang.org/x/crypto/bcrypt"
	"server/common"
	"strings"
)

func PromiseToString(promise int) string {
	switch promise {
	case common.USER:
		return "user"
	case common.EMPLOYEE:
		return "employee"
	case common.ADMIN:
		return "admin"
	default:
		return "myerr"
	}
}

func PromiseToInt(promise string) int {
	promise = strings.ToLower(promise)
	switch promise {
	case "user":
		return common.USER
	case "employee":
		return common.EMPLOYEE
	case "admin":
		return common.ADMIN
	default:
		return 0
	}
}

//func MD5(v string) string {
//	d := []byte(v)
//	m := md5.New()
//	m.Write(d)
//	return hex.EncodeToString(m.Sum(nil))
//}

// GetPwd 给密码加密
func GetPwd(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

func PwdOK(origin, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(origin), []byte(password)) == nil
}
