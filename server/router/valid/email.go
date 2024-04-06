package valid

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func EmailValid(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	return EmailValidString(email)
}

func EmailValidString(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*\\.[a-zA-Z0-9]{2,6}$")
	return re.MatchString(email)
}
