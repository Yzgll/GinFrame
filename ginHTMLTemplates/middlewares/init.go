package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义中间件
func InitMiddleware(c *gin.Context) {
	//判断用户是否登录
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
	//中间件之间和控制前路由之间共享数据
	//用context.Set设置，get读取
	c.Set("username", "张三")

	//定义一个协程统计日志
	cCp := c.Copy() //必须使用只读的副本
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Done! in path" + cCp.Request.URL.Path)
	}()
}
