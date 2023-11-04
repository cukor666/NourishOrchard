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
	err = db.Select("id", "name", "gender", "birthday", "phone", "address", "CreatedAt", "UpdatedAt").
		Where("promise = 1").
		Find(&result).Error
	return
}

// 根据用户ID查询用户
func (ud UserDao) SelectById(id uint) (user moudels.User, err error) {
	err = db.Select("id", "name", "gender", "birthday", "phone", "address", "CreatedAt", "UpdatedAt").
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
	err := db.Select("id", "name", "gender", "promise", "birthday", "phone", "address").
		Where("name = ?", name).
		Take(&user).Error
	if err != nil {
		log.Printf("select failed, err: %v\n", err)
		ok = false
		return
	}
	return
}

// 根据指定字段查询用户
func (ud UserDao) SelectByStruct(user moudels.User) (users []moudels.User, err error) {
	err = db.Select("id", "name", "gender", "birthday", "phone", "address").
		Where(&user).Take(&users).Error
	return
}

// 根据指定字段查询用户，携带有生日
func (ud UserDao) SelectByStructWithBirthday(user moudels.User, birthday []string) (users []moudels.User, err error) {
	err = db.Select("id", "name", "gender", "birthday", "phone", "address").
		Where(&user).Where("birthday between ? and ?", birthday[0], birthday[1]).
		Take(&users).Error
	return
}

// 更新用户信息
func (ud UserDao) UpdateUserInfo(user moudels.User) (ok bool) {
	ok = true
	err := db.Model(&user).Select("gender", "birthday", "phone", "address").Where("name = ?", user.Name).Updates(user).Error
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
