package api

import "github.com/gin-gonic/gin"

type authInfo struct {
	Username string `json:"username" binding:"required,lt=50"`
	Password string `json:"username" binding:"required,lt=50"`
}

func authHandler(c *gin.Context) {
	var auth authInfo
	if err := c.ShouldBind(&auth); err != nil {

	}


}
