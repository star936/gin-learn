package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"github_url": "https://www.github.com",
		"baidu_url": "https://www.baidu.com",
	})
}
