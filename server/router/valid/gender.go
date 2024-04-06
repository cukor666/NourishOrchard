package valid

import (
	"github.com/go-playground/validator/v10"
)

func GenderValid(fl validator.FieldLevel) bool {
	gender := fl.Field().String()
	return GenderValidString(gender)
}

func GenderValidString(gender string) bool {
	switch gender {
	case "男", "女":
		return true
	}
	return false
}
