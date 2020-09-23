package routers

import (
	"backend/middlewares"
	"backend/routers/api"
	"backend/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() {
	router := gin.New()
	router.Use(middlewares.Cors()) // Cors
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.GET("/api/ping", api.Ping)

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/users/:id", v1.GetUser)
		apiv1.GET("/users", v1.GetUsers)
		apiv1.POST("/users", v1.CreateUser)
		apiv1.PUT("/users/:id", v1.UpdateUser)
	}

	_ = router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
