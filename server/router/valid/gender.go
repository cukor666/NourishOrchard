package valid

import (
	"github.com/go-playground/validator/v10"
)

var GenderValid validator.Func = func(fl validator.FieldLevel) bool {
	gender := fl.Field().Interface().(string)
	switch gender {
	case "男", "女":
		return true
	}
	return false
}
