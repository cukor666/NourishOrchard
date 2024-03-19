package service

import (
	"errors"
	"regexp"
	"server/common/simpletool"
	"server/dao"
	"server/models"
	"server/request"
	"server/utils"
)

// Info 根据账号获取员工信息
func (e EmployeeService) Info(username string) (emp models.Employee, err error) {
	r := regexp.MustCompile("CZKJ[0-9]{3,}")
	matchString := r.MatchString(username)
	if !matchString {
		return models.Employee{}, errors.New("账号不符合系统规范")
	}
	return dao.EmployeeDao{}.SelectByUsername(username)
}

// Update 更新员工信息
func (e EmployeeService) Update(employee models.Employee) error {
	_, err := dao.EmployeeDao{}.Update(employee)
	return err
}

// ListWithPage 查询员工信息，带有分页
func (e EmployeeService) ListWithPage(p simpletool.Page, employee models.Employee) ([]models.Employee, int64, error) {
	result, total, err := dao.EmployeeDao{}.ListWithPage(p, employee)
	if err != nil {
		levelLog("查询失败")
		return nil, 0, err
	}
	return result, total, nil
}

// ForgetPassword 忘记密码
func (e EmployeeService) ForgetPassword(req request.ForgetPwdReq) error {
	_, err := dao.EmployeeDao{}.SelectByUsernameAndPhone(req.Username, req.Phone)
	if err != nil {
		levelLog("查询员工失败")
		return err
	}
	// 对新密码加密
	pwd, err := utils.GetPwd(req.Password)
	if err != nil {
		levelLog("新密码加密失败")
		return err
	}
	err = dao.AccountDao{}.ChangePassword(req.Username, string(pwd))
	if err != nil {
		levelLog("修改密码错误")
		return err
	}
	return nil
}
