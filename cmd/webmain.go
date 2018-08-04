package main

import (
	"github.com/gin-gonic/gin"
	"go-gin/pkg"
)

func main() {
	//创建一个路由handler
	router := gin.Default()
	//定义web模板文件路径
	router.LoadHTMLGlob("templates/*")

	router.GET("/upload", pkg.DemoWithWebForm1)

	router.Run(":8080")
}
