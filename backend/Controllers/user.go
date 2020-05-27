package Controllers

import (
	"backend/Services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"regexp"
)

var pattern *regexp.Regexp

func init() {
	pattern, _ = regexp.Compile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
}

// 用户信息查询 godoc
// @Summary 用户信息查询
// @Description 根据用户id查询用户信息
// @ID get-string-by-int
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} Resp
// @Failure 400 {string} Resp
// @router /api/users/{id} [get]
func GetUser(c *gin.Context) {
	var u Services.UserId
	if err := c.ShouldBindUri(&u); err != nil {
		BadRequest(c, err.Error())
		return
	}
	resp, err := Services.GetUser(u.Id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			NotFound(c, "Not found")
			return
		}
		RespHandler(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	RespHandler(c, http.StatusOK, "", resp)
}

// 用户集合查询 godoc
// @Summary 用户集合查询
// @Description 用户分页查询
// @Tags user
// @Accept json
// @Produce json
// @Param page query int true "page" mininum(1)
// @Param page_size query int true "page size" maxinum(100)
// @Success 200 {object} Resp
// @Failure 400 {string} Resp
// @router /api/users [get]
func GetUsers(c *gin.Context) {
	var u Services.UserPagination
	if err := c.ShouldBind(&u); err != nil {
		BadRequest(c, err.Error())
		return
	}
	resp, err := Services.GetUserPagination(&u)
	if err != nil {
		BadRequest(c, err.Error())
		return
	}
	RespHandler(c, http.StatusOK, "", resp)
}

// 用户注册 godoc
// @Summary 用户注册
// @Description 博客用户注册接口
// @Tags user
// @Accept json
// @Produce json
// @Param user body Services.User true "user info"
// @Success 200 {object} Resp
// @Failure 400 {string} Resp
// @router /api/users [post]
func UserRegister(c *gin.Context) {
	var user Services.User

	if err := c.ShouldBindJSON(&user); err != nil {
		BadRequest(c, err.Error())
		return
	}
	if !pattern.MatchString(user.Email) {
		BadRequest(c, "Email format is incorrect")
		return
	}

	result, err := Services.RecordNotExist(map[string]interface{}{"username": user.Username})
	if err != nil {
		RespHandler(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if !result {
		BadRequest(c, "Username exist")
		return
	}

	result, err = Services.RecordNotExist(map[string]interface{}{"email": user.Email})
	if err != nil {
		RespHandler(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if !result {
		BadRequest(c, "Email exist")
		return
	}

	resp, err := Services.CreateUser(&user)
	if err != nil {
		RespHandler(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	RespHandler(c, http.StatusOK, "", resp)
}
