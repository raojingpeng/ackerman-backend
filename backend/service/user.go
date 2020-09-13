package service

import (
	"backend/Middlewares"
	"backend/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserIdStruct struct {
	Id int `json:"id" binding:"required,gt=0"`
}

type CreateUserStruct struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type UserResp struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ExistUserById(id int) (bool, error) {
	return models.ExistUserById(id)
}

func GetUser(id int) (*UserResp, error) {
	user, err := models.GetUser(id)
	if err != nil {
		return nil, err
	}

	return &UserResp{
		Id:    user.ID,
		Name:  user.UserName,
		Email: user.Email,
	}, nil
}

func GetUsers(page, pageSize int) ([]*UserResp, error) {
	users, err := models.GetUsers(page, pageSize)
	if err != nil {
		return nil, err
	}

	var userResp []*UserResp
	for _, v := range users {
		userResp = append(userResp, &UserResp{
			Id:    v.ID,
			Name:  v.UserName,
			Email: v.Email,
		})
	}
	return userResp, nil
}

func (service *CreateUserStruct) Create() (*UserResp, error) {
	pwdHash := hashAndSalt([]byte(service.Password))

	var user = models.User{
		UserName:     service.UserName,
		Email:        service.Email,
		NickName:     service.NickName,
		PasswordHash: pwdHash,
	}
	if err := user.Create(); err != nil {
		return nil, err
	}

	return &UserResp{
		Id:    user.ID,
		Name:  user.UserName,
		Email: user.Email,
	}, nil
}

func UserExist(query map[string]interface{}) bool {
	var user models.User
	if err := db.Where(query).Find(user).Error; gorm.IsRecordNotFoundError(err) {
		return false
	}
	return true
}

type UpdateUserParams struct {
	Id    int    `uri:"id" binding:"required"`
	Name  string `json:"name" binding:"-"`
	Email string `json:"email" binding:"omitempty,email"`
}

func UpdateUser(params *UpdateUserParams) error {
	var user models.User
	if err := db.First(&user, params.Id).Error; err != nil {
		return err
	}

	err := db.Model(&user).Updates(&models.User{
		Name:  params.Name,
		Email: params.Email,
	}).Error

	return err
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
