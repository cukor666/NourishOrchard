package valid

import (
	"github.com/go-playground/validator/v10"
	mc "server/models/code"
	"server/utils/promisetool"
)

var PromiseValid validator.Func = func(fl validator.FieldLevel) bool {
	promise := fl.Field().Interface().(string)
	p := promisetool.ToInt(promise)
	switch p {
	case mc.USER, mc.EMPLOYEE, mc.ADMIN:
		return true
	}
	return false
}
