package app

import (
	"backend/pkg/e"
	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func RespHandler(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, Resp{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
}
