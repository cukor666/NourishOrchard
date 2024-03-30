package valid

import (
	"github.com/go-playground/validator/v10"
	"time"
)

func BirthdayValid(fl validator.FieldLevel) bool {
	birthday := fl.Field().Interface().(time.Time)
	return BirthdayValidTime(birthday)
}

func BirthdayValidTime(birthday time.Time) bool {
	age := time.Now().Year() - birthday.Year()
	return age >= 18
}
