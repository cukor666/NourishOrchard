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

func (ud UserDao) SelectAll() (result []moudels.User, err error) {
	err = db.Find(&result).Error
	return
}

func (ud UserDao) SelectById(id uint) (user moudels.User, err error) {
	err = db.Where("id = ?", id).Take(&user).Error
	return
}

func (ud UserDao) SelectUserByNameAndPwd(name, password string) (user moudels.User, ok bool) {
	ok = true
	err := db.Where("name = ? AND password = ?", name, password).Take(&user).Error
	if err != nil {
		log.Printf("select failed, err: %v\n", err)
		ok = false
		return
	}
	return
}

func (ud UserDao) SelectByName(name string) (user moudels.User, ok bool) {
	ok = true
	err := db.Where("name = ?", name).Take(&user).Error
	if err != nil {
		log.Printf("select failed, err: %v\n", err)
		ok = false
		return
	}
	return
}

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
