package valid

import (
	"fmt"
	"log"
	"regexp"
)

func levelLog(w string) {
	log.Println("valid层：", w)
}

// ValidUsername 校验用户名
func ValidUsername(username string) bool {
	matched, err := regexp.MatchString("^CZKJ\\d+$", username)
	if err != nil {
		levelLog(fmt.Sprintf("正则表达式使用错误"))
		return false
	}
	return matched
}
