package valid

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var PhoneValid validator.Func = func(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	compile := regexp.MustCompile("^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$")
	return compile.MatchString(phone)
}
