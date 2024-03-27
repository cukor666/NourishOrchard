package promisetool

import (
	mc "server/models/code"
	"strings"
)

func ToString(promise int) string {
	switch promise {
	case mc.USER:
		return "user"
	case mc.EMPLOYEE:
		return "employee"
	case mc.ADMIN:
		return "admin"
	}
	return ""
}

func ToInt(promise string) int {
	promise = strings.ToLower(promise)
	switch promise {
	case "user":
		return mc.USER
	case "employee":
		return mc.EMPLOYEE
	case "admin":
		return mc.ADMIN
	}
	return 0
}
