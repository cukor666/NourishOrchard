package main

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"nourish-orchard-internal/dao"
	"nourish-orchard-internal/models"
	"nourish-orchard-internal/utils"
	"regexp"
	"strconv"
)

const (
	EMPLOYEE = 3705
	ADMIN    = 8092
)

const (
	ADMIN_ID       = 0
	ADMIN_NAME     = 1
	ADMIN_PASSWORD = 2
	ADMIN_EMAIL    = 3
)

const (
	EMPLOYEE_ID       = 0
	EMPLOYEE_NAME     = 1
	EMPLOYEE_PASSWORD = 2
	EMPLOYEE_PHONE    = 3
	EMPLOYEE_POSITION = 4
	EMPLOYEE_SALARY   = 5
)

func main() {
	//f, err := excelize.OpenFile("excel/admin.xlsx")
	f, err := excelize.OpenFile("excel/employee.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	//insertAdmin(f, "Sheet1")
	insertEmployee(f, "Sheet1")
}

func insertAdmin(f *excelize.File, sheet string) {
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 去掉表头
	rows = rows[1:]
	var account models.Account
	var admin models.Admin
	account.Promise = ADMIN
	var registerDao dao.RegisterDao
	// 遍历数据
	for _, row := range rows {
		id, _ := strconv.Atoi(row[ADMIN_ID])
		admin.ID = uint(id)
		admin.Name = row[ADMIN_NAME]                     // 获取姓名
		password, ok := proPassword(row[ADMIN_PASSWORD]) // 获取密码
		if !ok {
			panic(errors.New(fmt.Sprintf("密码加盐失败，name: %s", admin.Name)))
		}
		account.Password = password // 加密之后的密码再存放到数据库中
		email := row[ADMIN_EMAIL]
		if ok = validEmail(email); !ok {
			panic(errors.New(fmt.Sprintf("邮箱校验失败，email: %s", email)))
		}
		admin.Email = email // 邮箱要经过校验之后再放到数据库中
		// 自动生成的属性
		username := utils.GenUsername()
		admin.Username = username
		account.Username = username
		uid, ok := registerDao.Register(account, admin)
		if !ok {
			fmt.Println("插入失败，name: ", admin.Name)
		} else {
			fmt.Println("插入成功，uid: ", uid)
		}
	}
}

func insertEmployee(f *excelize.File, sheet string) {
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 去掉表头
	rows = rows[1:]
	var account models.Account
	var employee models.Employee
	account.Promise = EMPLOYEE
	var registerDao dao.RegisterDao
	// 遍历数据
	for _, row := range rows {
		id, _ := strconv.Atoi(row[EMPLOYEE_ID])
		employee.ID = uint(id)
		employee.Name = row[EMPLOYEE_NAME]                  // 获取姓名
		password, ok := proPassword(row[EMPLOYEE_PASSWORD]) // 获取密码
		if !ok {
			panic(errors.New(fmt.Sprintf("密码加盐失败，name: %s", employee.Name)))
		}
		account.Password = password  // 加密之后的密码再存放到数据库中
		phone := row[EMPLOYEE_PHONE] // 获取电话
		ok = validPhone(phone)
		if !ok {
			panic(errors.New(fmt.Sprintf("电话号码校验失败，name: %s", employee.Name)))
		}
		employee.Phone = phone

		position := row[EMPLOYEE_POSITION]
		ok = validPosition(position)
		if !ok {
			panic(errors.New(fmt.Sprintf("职位校验失败，name: %s", employee.Name)))
		}
		salary, _ := strconv.Atoi(row[EMPLOYEE_SALARY])
		ok = validSalary(salary)
		if !ok {
			panic(errors.New(fmt.Sprintf("工资校验失败，name: %s", employee.Name)))
		}
		employee.Salary = salary

		// 自动生成的属性
		username := utils.GenUsername()
		employee.Username = username
		account.Username = username
		uid, ok := registerDao.RegisterEmployee(account, employee)
		if !ok {
			fmt.Println("插入失败，name: ", employee.Name)
		} else {
			fmt.Println("插入成功，uid: ", uid)
		}
	}
}

// 对密码进行加密，再返回加密后的密码
func proPassword(password string) (string, bool) {
	pwd, err := utils.GetPwd(password)
	if err != nil {
		return "", false
	}
	return string(pwd), true
}

// 邮箱校验
func validEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

// 电话验证
func validPhone(phone string) bool {
	pattern := `^\d{10,12}$`
	matched, _ := regexp.MatchString(pattern, phone)
	return matched
}

// 校验职位
func validPosition(position string) bool {
	return position == "采购员" || position == "仓库管理员"
}

// 校验工资
func validSalary(salary int) bool {
	return salary >= 3000 && salary <= 10000
}
