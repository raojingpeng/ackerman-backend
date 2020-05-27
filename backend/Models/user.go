package Models

import (
	"backend/Databases"
	"github.com/jinzhu/gorm"
	"go/types"
	"time"
)

type User struct {
	gorm.Model
	Username        string     `gorm:"column:username;size:64;unique_index"`
	Email           string     `gorm:"column:email;type:varchar(120);unique_index"`
	PasswordHash    string     `gorm:"column:password_hash;type:varchar(128)"`
	Token           string     `gorm:"column:token;size:32"`
	TokenExpiration *time.Time `gorm:"column:token_expiration"`
}

type Users []User

func init() {
	if !Databases.Mysql.HasTable(&User{}) {
		if err := Databases.Mysql.CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
}

func (User) TableName() string {
	return "user"
}

func (u *Users) Pagination(page int, pageSize int) (err error) {
	err = Databases.Mysql.Offset((page - 1) * pageSize).Limit(pageSize).Find(u).Error
	return
}

func (u *User) QueryById(id int) (err error) {
	err = Databases.Mysql.First(u, id).Error
	return
}

func (u *User) QueryByFirst(query interface{}, args ...interface{}) (err error) {
	var db *gorm.DB

	switch query.(type) {
	case types.Nil:
		db = Databases.Mysql
	default:
		db = Databases.Mysql.Where(query, args...)
	}

	err = db.First(u).Error
	return
}

func (u *Users) QueryByFind(query interface{}, args ...interface{}) (err error) {
	var db *gorm.DB

	switch query.(type) {
	case types.Nil:
		db = Databases.Mysql
	default:
		db = Databases.Mysql.Where(query, args...)
	}

	err = db.Find(u).Error
	return err
}

func (u *User) Insert() error {
	err := Databases.Mysql.Create(u).Error
	return err
}

func (u *User) Update(column string, value interface{}) error {
	err := Databases.Mysql.Model(u).Update(column, value).Error
	return err
}
