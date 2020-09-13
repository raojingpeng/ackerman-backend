package v1

import (
	"backend/pkg/app"
	"backend/pkg/e"
	"backend/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUser godoc
// @Summary 用户查询
// @Description 根据用户id查询用户信息
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} app.Resp
// @Failure 400 {object} app.Resp
// @Router /api/v1/users/{id} [get]
func GetUser(c *gin.Context) {
	var param service.UserIdStruct
	if err := c.ShouldBindUri(&param); err != nil {
		app.RespHandler(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	exist, err := service.ExistUserById(param.Id)
	if err != nil {
		app.RespHandler(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, nil)
		return
	}
	if !exist {
		app.RespHandler(c, http.StatusOK, e.ERROR_NOT_EXIST_USER, nil)
		return
	}

	user, err := service.GetUser(param.Id)
	if err != nil {
		app.RespHandler(c, http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}
	app.RespHandler(c, http.StatusOK, e.SUCCESS, user)
}

type UserPagination struct {
	Page     int `form:"page" binding:"required,gte=1"`
	PageSize int `form:"page_size" binding:"required,gte=5,lte=100"`
}

// GetUsers godoc
// @Summary 用户分页查询
// @Description 用户分页查询
// @Tags user
// @Accept json
// @Produce json
// @Param page query int true "page" mininum(1)
// @Param page_size query int true "page size" maxinum(100)
// @Success 200 {object} app.Resp
// @Failure 400 {object} app.Resp
// @router /api/users [get]
func GetUsers(c *gin.Context) {
	var param UserPagination
	if err := c.ShouldBind(&param); err != nil {
		app.RespHandler(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	users, err := service.GetUsers(param.Page, param.PageSize)
	if err != nil {
		app.RespHandler(c, http.StatusInternalServerError, e.ERROR_GET_USERS_FAIL, nil)
		return
	}
	app.RespHandler(c, http.StatusOK, e.SUCCESS, users)
}

func checkUser(username string, email string) error {
	if username == "" && email == "" {
		return errors.New("username and password can not be empty")
	}
	if username != "" && service.UserExist(map[string]interface{}{"name": username}) {
		return errors.New("username exist")
	}
	if email != "" && service.UserExist(map[string]interface{}{"email": email}) {
		return errors.New("email exist")
	}
	return nil
}

// @Summary 用户注册
// @Description 新用户注册
// @Tags user
// @Accept json
// @Produce json
// @Param user body service.CreateUserParams true "user info"
// @Success 200 {object} Resp
// @Failure 400 {object} Resp
// @router /api/users [post]
func CreateUser(c *gin.Context) {
	var u service.CreateUserParams

	if err := c.ShouldBindJSON(&u); err != nil {
		BadRequest(c, err.Error())
		return
	}

	if err := checkUser(u.Name, u.Email); err != nil {
		BadRequest(c, err.Error())
	}

	resp, err := service.CreateUser(&u)

	if err != nil {
		RespHandler(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	RespHandler(c, http.StatusOK, "", resp)
}

// @Summary 更新用户信息
// @Description 博客用户更新接口
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Param user body service.UpdateUserParams true "user update form"
// @Success 200 {object} Resp
// @Failure 400 {object} Resp
// @router /api/users/{id} [put]
func UpdateUser(c *gin.Context) {
	var u service.UpdateUserParams

	if err := c.ShouldBindJSON(&u); err != nil {
		BadRequest(c, err.Error())
		return
	}

	if err := checkUser(u.Name, u.Email); err != nil {
		BadRequest(c, err.Error())
		return
	}

	if err := service.UpdateUser(&u); err != nil {
		BadRequest(c, err.Error())
		return
	}

	RespHandler(c, http.StatusOK, "", nil)
}
