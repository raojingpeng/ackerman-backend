package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name            string     `gorm:"column:name;size:64;unique_index"`
	Email           string     `gorm:"column:email;type:varchar(120);unique_index"`
	PasswordHash    string     `gorm:"column:password_hash;type:varchar(128)"`
	Token           string     `gorm:"column:token;size:32"`
	TokenExpiration *time.Time `gorm:"column:token_expiration"`
}

func init() {
	if !db.HasTable(&User{}) {
		if err := db.CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
}

func ExistUserById(id int) (bool, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
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
	err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}

func CreateUser(data map[string]interface{}) error {
	user := User{
		Name:         data["name"].(string),
		Email:        data["email"].(string),
		PasswordHash: data["password_hash"].(string),
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
