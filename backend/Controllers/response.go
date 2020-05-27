package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	Status int         `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

func BadRequest(c *gin.Context, errMsg string) {
	c.JSON(http.StatusBadRequest, Resp{
		Status: http.StatusBadRequest,
		Error:  errMsg,
		Data:   nil,
	})
}

func NotFound(c *gin.Context, errMsg string) {
	c.JSON(http.StatusNotFound, Resp{
		Status: http.StatusNotFound,
		Error:  errMsg,
		Data:   nil,
	})
}

func RespHandler(c *gin.Context, errCode int, errMsg string, data interface{}) {
	c.JSON(errCode, Resp{
		Status: errCode,
		Error:  errMsg,
		Data:   data,
	})
}
