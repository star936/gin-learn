package main

import (
	"gin-learn/routers"
)

func main() {
	router := routers.InitRouter()
	_ = router.Run(":5000")
}
