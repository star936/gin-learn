package apis

import (
	"gin-learn/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGitHubUserInfo(ctx *gin.Context)  {
	code := ctx.Query("code")
	accessToken := pkg.GHO.GetAccessToken(code)
	ctx.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}

func GetBaiduUserInfo(ctx *gin.Context)  {
	code := ctx.Query("code")
	accessToken := pkg.BDO.GetAccessToken(code)
	ctx.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
