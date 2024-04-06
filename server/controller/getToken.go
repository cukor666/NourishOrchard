package controller

import (
	"errors"
	"fmt"
	"server/common/levellog"
	"strings"
)

// GetToken 从authorization中获取token
func GetToken(authorization string) (token string, err error) {
	// 获取token
	split := strings.Split(authorization, " ")
	if len(split) != 2 {
		levellog.Controller("获取token失败")
		return "", errors.New(fmt.Sprintf("获取token失败"))
	}
	return split[1], nil
}
