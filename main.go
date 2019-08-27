package main

import (
	"gin-learn/config"
	"gin-learn/routers"
)

func init() {
	config.SetUp()
}

func main() {
	router := routers.InitRouter()
	_ = router.Run(":5000")
}
