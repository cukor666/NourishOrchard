package spliterr

import (
	"regexp"
)

func GetErrMsg(str string) string {
	re := regexp.MustCompile("Error:")
	split := re.Split(str, -1)
	if len(split) > 1 {
		return split[1]
	}
	return str
}
