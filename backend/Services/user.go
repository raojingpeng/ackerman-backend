package Services

import (
	"backend/Middlewares"
	"backend/Models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserId struct {
	Id int `uri:"id" binding:"required"`
}

type UserPagination struct {
	Page     int `form:"page" binding:"required,gte=1"`
	PageSize int `form:"page_size" binding:"required,gte=5,lte=100"`
}

type User struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResp struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GetUser(id int) (*UserResp, error) {
	var user Models.User
	err := user.QueryById(id)
	if err != nil {
		return nil, err
	}
	return &UserResp{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func CreateUser(u *User) (*UserResp, error) {
	pwdHash := hashAndSalt([]byte(u.Password))

	var user Models.User
	user.Username = u.Username
	user.Email = u.Email
	user.PasswordHash = pwdHash

	if err := user.Insert(); err != nil {
		return nil, err
	}
	return &UserResp{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func GetUserPagination(u *UserPagination) (*[]UserResp, error) {
 	var users Models.Users
	err := users.Pagination(u.Page, u.PageSize)
	var userResp []UserResp
	for _, v := range users {
		userResp = append(userResp, UserResp{
			Id:       v.ID,
			Username: v.Username,
			Email:    v.Email,
		})
	}
	return &userResp, err
}

func RecordNotExist(m map[string]interface{}) (bool, error) {
	var user Models.User
	err := user.QueryByFirst(m)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return true, nil
		}
		return false, err
	}
	return false, nil
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
