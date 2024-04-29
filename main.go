package main

import (
	"fmt"
	"k8s.io/klog/v2"
	"operator/config"
	route "operator/router"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	route.InitRoute(engine)
	if err := engine.Run(fmt.Sprintf("%s:%s", config.GetString(config.ServerHost), config.GetString(config.ServerPort))); err != nil {
		klog.Fatal(err)
		return
	}
}
