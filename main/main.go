package main

import (
	"gin"
	"net/http"
)

func main() {
	// 创建一个 gin Engine，本质上是一个 http Handler
	mux := gin.Default()
	// 注册中间件
	mux.Use()
	// 注册一个 path 为 /ping 的处理函数
	mux.POST("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pone")
	})
	// 运行 http 服务
	if err := mux.Run(":8080"); err != nil {
		panic(err)
	}
}
