package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET",
	})
}
func main() {
	router := gin.Default() //获取默认路由
	router.GET("/ping", Show)
	router.DELETE("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "这是一个DELETE请求，主要用于从服务器删除资源")
	})
	router.PUT("/ping", func(ctx *gin.Context) {
		ctx.String(200, "这是一个PUT请求，主要用于从服务器更新资源")
	})
	router.POST("/ping", func(ctx *gin.Context) {
		ctx.String(200, "这是一个POST请求，主要用于从服务器创建资源")
	})

	router.Run()
}
