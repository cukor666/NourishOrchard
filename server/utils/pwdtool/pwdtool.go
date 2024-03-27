package pwdtool

import "golang.org/x/crypto/bcrypt"

// GetPwd 给密码加密
func GetPwd(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

// PwdOK 校验密码是否正确
func PwdOK(origin, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(origin), []byte(password)) == nil
}
