package dao

import (
	"Gin/app"
	"Gin/models"
	"Gin/vo"
	"log"
)

type UserDao struct{}

// Insert 插入用户数据
func (ud UserDao) Insert(user models.User) (newId uint, affect int64) {
	d := app.MySQLDB().Create(&user)
	if d.Error != nil {
		log.Printf("insert failed")
		return 0, 0
	}
	return user.ID, d.RowsAffected
}

// SelectAll 查询所有用户，包括普通用户、管理员、超级管理员
func (ud UserDao) SelectAll() (result []models.User, err error) {
	err = app.MySQLDB().Find(&result).Error
	return
}

// SelectAllUser 查询所有用户，权限为1的
func (ud UserDao) SelectAllUser() (result []models.User, err error) {
	// 不把密码查出来给前端，防止密码泄露
	err = app.MySQLDB().Omit("password", "DeletedAt").Where("promise = ?", 1).Find(&result).Error
	return
}

// SelectAllUserByLimit 查询所有普通用户，分页查询
func (ud UserDao) SelectAllUserByLimit(pageSize, currentPage int) (users []models.User, result vo.Result) {
	// 不把密码查出来给前端，防止密码泄露
	tx := app.MySQLDB().Omit("password", "DeletedAt").Where("promise = ?", 1)
	result.Rows = tx.Find(&[]models.User{}).RowsAffected
	result.Error = tx.Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&users).Error
	return
}

// SelectById 根据用户ID查询用户
func (ud UserDao) SelectById(id uint) (user models.User, err error) {
	err = app.MySQLDB().Omit("password", "DeletedAt").Where("id = ?", id).Take(&user).Error
	return
}

// SelectUserByNameAndPwd 通过用户名和密码查询用户
func (ud UserDao) SelectUserByNameAndPwd(name, password string) (user models.User, ok bool) {
	err := app.MySQLDB().
		Where("name = ? AND password = ?", name, password).
		Take(&user).Error
	if err != nil {
		log.Printf("select failed, err: %v\n", err)
		return models.User{}, false
	}
	return user, true
}

// SelectByName 根据用户名查询个人信息
func (ud UserDao) SelectByName(name string) (user models.User, ok bool) {
	// 指定前端要看到的字段，包含隐私信息
	err := app.MySQLDB().Omit("password", "DeletedAt").
		Where("name = ?", name).
		Take(&user).Error
	if err != nil {
		log.Printf("select failed, err: %v\n", err)
		return models.User{}, false
	}
	return user, true
}

// SelectByStruct 根据指定字段查询用户
func (ud UserDao) SelectByStruct(user models.User, currentPage, pageSize int) (users []models.User, rows int64, err error) {
	// name、nick_name、phone、address等字段需要模糊查询才能有意义，所以得单独提取出来
	name, phone, address, nickname := user.Name, user.Phone, user.Address, user.NickName
	user.Name, user.Phone, user.Address, user.NickName = "", "", "", ""
	tx := app.MySQLDB().Omit("password", "DeletedAt").Where(&user).
		Where("name LIKE ?", "%"+name+"%").
		Where("phone LIKE ?", "%"+phone+"%").
		Where("address LIKE ?", "%"+address+"%").
		Where("nick_name LIKE ?", "%"+nickname+"%")
	rows = tx.Find(&[]models.User{}).RowsAffected
	err = tx.Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&users).Error
	return
}

// SelectByStructWithBirthday 根据指定字段查询用户，携带有生日
func (ud UserDao) SelectByStructWithBirthday(user models.User, currentPage, pageSize int, birthday []string) (users []models.User, rows int64, err error) {
	// phone、address等字段需要模糊查询才能有意义，所以得单独提取出来
	name, phone, address, nickname := user.Name, user.Phone, user.Address, user.NickName
	user.Name, user.Phone, user.Address, user.NickName = "", "", "", ""
	tx := app.MySQLDB().Omit("password", "DeletedAt").Where(&user).
		Where("name LIKE ?", "%"+name+"%").
		Where("phone LIKE ?", "%"+phone+"%").
		Where("address LIKE ?", "%"+address+"%").
		Where("nick_name LIKE ?", "%"+nickname+"%").
		Where("birthday between ? and ?", birthday[0], birthday[1])
	rows = tx.Find(&[]models.User{}).RowsAffected
	err = tx.Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&users).Error
	return
}

// UpdateUserInfo 更新用户信息
func (ud UserDao) UpdateUserInfo(user models.User) (ok bool) {
	var err error
	if len(user.NickName) != 0 { // 如果外部传了nickname后端才修改,否则不做修改
		err = app.MySQLDB().Model(&user).
			Omit("ID", "name", "password", "promise", "CreatedAt", "DeletedAt").
			Where("name = ?", user.Name).
			Updates(user).Error
	} else {
		err = app.MySQLDB().Model(&user).
			Omit("ID", "name", "password", "promise", "CreatedAt", "DeletedAt", "nick_name").
			Where("name = ?", user.Name).
			Updates(user).Error
	}
	if err != nil {
		log.Printf("update failed, err: %v\n", err)
		return false
	}
	return true
}

// DeleteById 通过ID删除用户
func (ud UserDao) DeleteById(id uint) (ok bool) {
	err := app.MySQLDB().Delete(&models.User{}, id).Error
	if err != nil {
		log.Printf("delete failed, err: %v\n", err)
		return false
	}
	return true
}

// SelectPromiseByName 通过用户名获取用户权限
func (ud UserDao) SelectPromiseByName(name string) (int8, bool) {
	var user models.User
	err := app.MySQLDB().Select("promise").Where("name = ?", name).Take(&user).Error
	if err != nil {
		log.Printf("select promise failed, err: %v\n", err)
		return -1, false
	}
	return user.Promise, true
}

// SelectUserCount 获取指定权限的用户个数
func (ud UserDao) SelectUserCount(promise int8) (userCount int64, err error) {
	err = app.MySQLDB().Model(&models.User{}).
		Where("promise = ?", promise).
		Count(&userCount).Error
	return
}
