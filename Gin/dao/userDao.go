package dao

import (
	"Gin/moudels"
	"log"
	"time"
)

type UserDao struct{}

func init() {
	db.AutoMigrate(&moudels.User{})
}

// 插入用户数据
func (ud UserDao) Insert(user moudels.User) (newId uint, affect int64) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	d := db.Create(&user)
	if d.Error != nil {
		log.Printf("insert failed")
		return 0, 0
	}
	return user.ID, d.RowsAffected
}

// 查询所有用户，包括普通用户、管理员、超级管理员
func (ud UserDao) SelectAll() (result []moudels.User, err error) {
	err = db.Find(&result).Error
	return
}

// 查询所有用户，权限为1的
func (ud UserDao) SelectAllUser() (result []moudels.User, err error) {
	// 不把密码查出来给前端，防止密码泄露
	err = db.Omit("password", "DeletedAt").
		Where("promise = ?", 1).
		Find(&result).Error
	return
}

// 根据用户ID查询用户
func (ud UserDao) SelectById(id uint) (user moudels.User, err error) {
	err = db.Omit("password", "DeletedAt").
		Where("id = ?", id).
		Take(&user).Error
	return
}

// 通过用户名和密码查询用户
func (ud UserDao) SelectUserByNameAndPwd(name, password string) (user moudels.User, ok bool) {
	ok = true
	err := db.Select("name", "password").
		Where("name = ? AND password = ?", name, password).
		Take(&user).Error
	if err != nil {
		log.Printf("select failed, err: %v\n", err)
		ok = false
		return
	}
	return
}

// 根据用户名查询个人信息
func (ud UserDao) SelectByName(name string) (user moudels.User, ok bool) {
	ok = true
	// 指定前端要看到的字段，包含隐私信息
	err := db.Omit("password", "DeletedAt").
		Where("name = ?", name).
		Take(&user).Error
	if err != nil {
		log.Printf("select failed, err: %v\n", err)
		ok = false
		return
	}
	user.Show()
	return
}

// 根据指定字段查询用户
func (ud UserDao) SelectByStruct(user moudels.User) (users []moudels.User, err error) {
	// phone、address等字段需要模糊查询才能有意义，所以得单独提取出来
	phone := user.Phone
	user.Phone = "" // 提取完之后清空user中的属性，防止下面的Where使用phone = 'xxx'，下address同
	address := user.Address
	user.Address = ""
	err = db.Omit("password", "DeletedAt").
		Where(&user).Where("phone LIKE ? AND address LIKE ?", "%"+phone+"%", "%"+address+"%").
		Find(&users).Error
	return
}

// 根据指定字段查询用户，携带有生日
func (ud UserDao) SelectByStructWithBirthday(user moudels.User, birthday []string) (users []moudels.User, err error) {
	// phone、address等字段需要模糊查询才能有意义，所以得单独提取出来
	phone := user.Phone
	user.Phone = "" // 提取完之后清空user中的属性，防止下面的Where使用phone = 'xxx'，下address同
	address := user.Address
	user.Address = ""
	err = db.Select("id", "name", "gender", "birthday", "phone", "address", "CreatedAt", "UpdatedAt").
		Where(&user).Where("birthday between ? and ?", birthday[0], birthday[1]).
		Where("phone LIKE ? AND address LIKE ?", "%"+phone+"%", "%"+address+"%").
		Find(&users).Error
	return
}

// 更新用户信息
func (ud UserDao) UpdateUserInfo(user moudels.User) (ok bool) {
	ok = true
	var err error
	if len(user.NickName) != 0 { // 如果外部传了nickname后端才修改,否则不做修改
		err = db.Model(&user).
			Select("*").
			Omit("ID", "name", "password", "promise", "CreatedAt", "DeletedAt").
			Where("name = ?", user.Name).
			Updates(user).Error
	} else {
		err = db.Model(&user).
			Select("*").
			Omit("ID", "name", "password", "promise", "CreatedAt", "DeletedAt", "nick_name").
			Where("name = ?", user.Name).
			Updates(user).Error
	}
	if err != nil {
		log.Printf("update failed, err: %v\n", err)
		ok = false
		return
	}
	return
}

// 通过ID删除用户
func (ud UserDao) DeleteById(id uint) (ok bool) {
	ok = true
	err := db.Delete(&moudels.User{}, id).Error
	if err != nil {
		log.Printf("delete failed, err: %v\n", err)
		ok = false
		return
	}
	return
}

// 通过用户名获取用户权限
func (ud UserDao) SelectPromiseByName(name string) (promise int8, ok bool) {
	ok = true
	var user moudels.User
	err := db.Select("promise").Where("name = ?", name).Take(&user).Error
	promise = user.Promise
	if err != nil {
		log.Printf("select promise failed, err: %v\n", err)
		ok = false
		return
	}
	return
}
