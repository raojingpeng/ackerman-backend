package Router

import (
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
		api.GET("/ping", ping)
	}

	_ = router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
