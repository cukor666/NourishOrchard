package valid

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"server/common/levellog"
)

var UsernameValid validator.Func = func(fl validator.FieldLevel) bool {
	username := fl.Field().Interface().(string)
	matched, err := regexp.MatchString("CZKJ\\d+", username)
	if err != nil {
		levellog.Valid(err.Error())
		return false
	}
	return matched
}
