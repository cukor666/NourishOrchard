package valid

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"server/common/levellog"
)

func UsernameValid(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	return UsernameValidString(username)
}

func UsernameValidString(username string) bool {
	matched, err := regexp.MatchString("CZKJ\\d+", username)
	if err != nil {
		levellog.Valid(err.Error())
		return false
	}
	return matched
}
