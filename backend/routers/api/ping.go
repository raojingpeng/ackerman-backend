package api

import "github.com/gin-gonic/gin"

// PingExample godoc
// @Summary ping example
// @Description Do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @routers /api/ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
