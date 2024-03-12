package service

import (
	"log"
)

type AccountService struct{}
type LoginService struct{}
type RegisterService struct{}
type UserService struct{}
type EmployeeService struct{}
type AdminService struct{}

func levelLog(w string) {
	log.Println("service层：", w)
}
