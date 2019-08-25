package routers

import (
	"gin-learn/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controllers.Index)
	return router
}
