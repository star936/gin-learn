package routers

import (
	"gin-learn/apis"
	"gin-learn/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controllers.Index)
	router.GET("/github/oauth/redirect", apis.GetGitHubUserInfo)
	router.GET("/baidu/oauth/redirect", apis.GetBaiduUserInfo)
	return router
}
