package apis

import (
	"gin-learn/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGitHubUserInfo(ctx *gin.Context)  {
	code := ctx.Query("code")
	accessToken := pkg.GHO.GetAccessToken(code)
	name := pkg.GHO.GetUserInfo(accessToken)
	ctx.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func GetBaiduUserInfo(ctx *gin.Context)  {
	code := ctx.Query("code")
	accessToken := pkg.BDO.GetAccessToken(code)
	name := pkg.BDO.GetUserInfo(accessToken)
	ctx.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}
