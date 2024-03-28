package valid

import (
	"github.com/go-playground/validator/v10"
	"time"
)

var BirthdayValid validator.Func = func(fl validator.FieldLevel) bool {
	birthday := fl.Field().Interface().(time.Time)
	age := time.Now().Year() - birthday.Year()
	return age >= 18
}
