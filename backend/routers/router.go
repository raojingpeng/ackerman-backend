package routers

import (
	"backend/controller"
	"backend/Middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRouter() {
	router := gin.New()
	router.Use(Middlewares.Cors()) // Cors
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	api := router.Group("/api")
	{
		api.GET("/ping", controller.Ping)
		api.GET("/users/:id", controller.GetUser)
		api.POST("/users", controller.UserRegister)
		api.GET("/users", controller.GetUsers)
		api.PUT("/users/:id", controller.UpdateUser)
	}

	_ = router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
