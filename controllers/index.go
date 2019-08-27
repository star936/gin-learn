package controllers

import (
	"gin-learn/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"github_url": pkg.GHO.GetAuthURL(),
		"baidu_url":  pkg.BDO.GetAuthURL(),
	})
}
