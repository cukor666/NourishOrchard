package valid

import "github.com/go-playground/validator/v10"

func EmpPosition(fl validator.FieldLevel) bool {
	ep := fl.Field().String()
	return EmpPositionString(ep)
}

func EmpPositionString(ep string) bool {
	switch ep {
	case "采购员", "仓库管理员":
		return true
	}
	return false
}
