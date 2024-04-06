package valid

import "github.com/go-playground/validator/v10"

func AddressValid(fl validator.FieldLevel) bool {
	addr := fl.Field().String()
	return AddressValidString(addr)
}

func AddressValidString(addr string) bool {
	addrLen := len(addr)
	if addrLen == 0 {
		return true
	}
	if addrLen < 3 || addrLen > 200 {
		return false
	}
	return true
}
