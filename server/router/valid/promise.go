package valid

import (
	"github.com/go-playground/validator/v10"
	mc "server/models/code"
	"server/utils/promisetool"
)

func PromiseValid(fl validator.FieldLevel) bool {
	promise := fl.Field().String()
	return PromiseValidString(promise)
}

func PromiseValidString(promise string) bool {
	p := promisetool.ToInt(promise)
	switch p {
	case mc.USER, mc.EMPLOYEE, mc.ADMIN:
		return true
	}
	return false
}
