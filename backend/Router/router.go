package Router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"backend/Middlewares"
)

func InitRouter() {
	router := gin.New()
	// 跨域中间件
	router.Use(Middlewares.Cors())

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	api := router.Group("/api")
	{
		api.GET("/ping", ping)
	}

	_ = router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
