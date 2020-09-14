package v1

import (
	"backend/pkg/app"
	"backend/pkg/e"
	"backend/service"
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
// @Failure 500 {object} app.Resp
// @Router /api/v1/users/{id} [get]
func GetUser(c *gin.Context) {
	var s service.UserIdStruct
	if err := c.ShouldBindUri(&s); err != nil {
		app.RespHandler(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	exist, err := service.ExistUser(map[string]interface{}{"id": s.Id})
	if err != nil {
		app.RespHandler(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER_FAIL, nil)
		return
	}
	if !exist {
		app.RespHandler(c, http.StatusOK, e.ERROR_NOT_EXIST_USER, nil)
		return
	}

	user, err := service.GetUser(s.Id)
	if err != nil {
		app.RespHandler(c, http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}
	app.RespHandler(c, http.StatusOK, e.SUCCESS, user)
}

// GetUsers godoc
// @Summary 用户分页查询
// @Description 用户分页查询接口
// @Tags user
// @Accept json
// @Produce json
// @Param page query int true "page" mininum(1)
// @Param page_size query int true "page size" maxinum(100)
// @Success 200 {object} app.Resp
// @Failure 500 {object} app.Resp
// @router /api/v1/users [get]
func GetUsers(c *gin.Context) {
	var s service.PageInfo
	if err := c.ShouldBind(&s); err != nil {
		app.RespHandler(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	resp, err := service.GetUsers(s.Page, s.PageSize)
	if err != nil {
		app.RespHandler(c, http.StatusInternalServerError, e.ERROR_GET_USERS_FAIL, nil)
		return
	}
	app.RespHandler(c, http.StatusOK, e.SUCCESS, resp)
}

// @Summary 用户注册
// @Description 新用户注册接口
// @Tags user
// @Accept json
// @Produce json
// @Param user body service.CreateUserStruct true "user info"
// @Success 200 {object} app.Resp
// @Failure 500 {object} app.Resp
// @router /api/v1/users [post]
func CreateUser(c *gin.Context) {
	var s service.CreateUserStruct

	if err := c.ShouldBindJSON(&s); err != nil {
		app.RespHandler(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	result, err := service.ExistUser(map[string]interface{}{"username": s.UserName})
	if err != nil {
		app.RespHandler(c, http.StatusInternalServerError, e.ERROR_QUERY_USER_FAIL, nil)
		return
	}
	if result {
		app.RespHandler(c, http.StatusOK, e.ERROR_EXIST_USERNAME, nil)
		return
	}

	resp, err := s.Create()
	if err != nil {
		app.RespHandler(c, http.StatusInternalServerError, e.ERROR_CREATE_USER_FAIL, nil)
		return
	}

	app.RespHandler(c, http.StatusOK, e.SUCCESS, resp)
}

// @Summary 用户信息更新
// @Description 用户信息更新接口
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Param user body service.UpdateUserStruct true "user update form"
// @Success 200 {object} app.Resp
// @Failure 500 {object} app.Resp
// @router /api/v1/users/{id} [put]
func UpdateUser(c *gin.Context) {
	var s service.UpdateUserStruct

	if err := c.ShouldBindJSON(&s); err != nil {
		app.RespHandler(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	result, err := service.ExistUser(map[string]interface{}{"username": s.UserName})
	if err != nil {
		app.RespHandler(c, http.StatusInternalServerError, e.ERROR_QUERY_USER_FAIL, nil)
		return
	}
	if result {
		app.RespHandler(c, http.StatusOK, e.ERROR_EXIST_USERNAME, nil)
		return
	}

	if err := s.Update(); err != nil {
		app.RespHandler(c, http.StatusInternalServerError, e.ERROR_UPDATE_USER_FAIL, nil)
		return
	}
	app.RespHandler(c, http.StatusOK, e.SUCCESS, nil)
}
