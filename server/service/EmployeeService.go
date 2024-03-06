package service

import (
	"errors"
	"regexp"
	"server/dao"
	"server/models"
)

// Info 根据账号获取员工信息
func (e EmployeeService) Info(username string) (emp models.Employee, err error) {
	r := regexp.MustCompile("CZKJ[0-9]{3,}")
	matchString := r.MatchString(username)
	if !matchString {
		return models.Employee{}, errors.New("账号不符合系统规范")
	}
	return employeeDao.SelectByUsername(username)
}

// Update 更新员工信息
func (e EmployeeService) Update(employee models.Employee) error {
	_, err := dao.EmployeeDao{}.Update(employee)
	return err
}
