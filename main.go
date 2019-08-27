package main

import (
	"gin-learn/config"
	"gin-learn/pkg"
	"gin-learn/routers"
)

func init() {
	config.SetUp()
	pkg.SetUp()
}

func main() {
	router := routers.InitRouter()
	_ = router.Run(":5000")
}
