package main

//中间件就是匹配路由之前和完成之后的一系列操作
//中间件中启动协程时不能直接使用原始的上下文
//必须使用它的c.Copy()的只能可读的副本
import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// c.Next()会暂停当前中间件的执行，去执行下一个中间件或主处理函数，然后再回来继续执行后续的代码。
//c.Abort表示终止调用这个请求,不再执行后续任何处理函数或中间件

func initMiddlewareOne(c *gin.Context) {

	fmt.Println("1-我是一个中间件-initMiddlewareOne")
	//调用该请求的剩余处理程序
	c.Next()
	fmt.Println("2-我是一个中间件-initMiddlewareOne")

}
func initMiddlewareTwo(c *gin.Context) {

	fmt.Println("1-我是一个中间件-initMiddlewareTwo")
	//调用该请求的剩余处理程序
	c.Next()
	fmt.Println("2-我是一个中间件-initMiddlewareTwo")

}
func main() {
	r := gin.Default()
	//配置模板
	// r.LoadHTMLGlob("templates/**/*")
	//配置静态路由
	//get的参数是一个路由和多个函数
	//GET(relativePath string, handlers ...HandlerFunc)
	// r.Static("/static", "./static")

	//路由后面可以有多个回调函数，最后一个回调函数之前的
	//回调函数都叫中间件

	//r.use配置多个全局中间件
	r.Use(initMiddlewareOne, initMiddlewareTwo)
	r.GET("/", func(c *gin.Context) {
		fmt.Println("这是一个首页")
		time.Sleep(time.Second * 5)
		c.String(200, "gin首页")
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(200, "新闻页面")
	})
	r.GET("/login", func(c *gin.Context) {
		c.String(200, "登录页面")
	})
	r.Run()
}
