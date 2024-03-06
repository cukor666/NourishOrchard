package service

import (
	"log"
	"server/dao"
)

type AccountService struct{}
type LoginService struct{}
type RegisterService struct{}
type UserService struct{}
type EmployeeService struct{}
type AdminService struct{}

var (
	accountDao  dao.AccountDao
	userDao     dao.UserDao
	employeeDao dao.EmployeeDao
	adminDao    dao.AdminDao
)

func levelLog(w string) {
	log.Println("service层：", w)
}
