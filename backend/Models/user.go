package Models

import (
	"backend/Databases"
	"backend/Middlewares"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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

func init() {
	// Migrate the schema
	Databases.DB.AutoMigrate(&User{})
}

func (User) TableName() string {
	return "user"
}

func (u *User) SetPassword(pwd string) {
	pwdHash := hashAndSalt([]byte(pwd))
	Databases.DB.Model(u).Update("password_hash", pwdHash)
}

func (u *User) CheckPassword(pwd string) bool {
	return comparePasswords(u.PasswordHash, []byte(pwd))
}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		Middlewares.Log.Printf(err.Error())
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}
