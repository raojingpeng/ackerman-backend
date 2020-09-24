package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName     string `gorm:"column:username;size:64;unique_index"`
	PasswordHash string `gorm:"column:password_hash;type:varchar(128)"`
	NickName     string `gorm:"column:nickname;size:64"`
	Avatar       string `gorm:"column:avatar;size:1000"`
	Email        string `gorm:"column:email;type:varchar(120);unique_index"`
}

func ExistUser(data interface{}) (bool, error) {
	if err := db.Where(data).First(&User{}).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func GetUser(id int) (*User, error) {
	var user User
	err := db.First(&user, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

func GetUsers(page, pageSize int) ([]*User, error) {
	var users []*User
	err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}

func (u *User) Create() error {
	if err := db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(id int, data interface{}) error {
	if err := db.First(&User{}, id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
