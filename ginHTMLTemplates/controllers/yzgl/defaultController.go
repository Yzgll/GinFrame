package yzgl

import (
	"fmt"
	"ginframe/ginHTMLTemplates/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func (def DefaultController) Index(c *gin.Context) {
	//设置session
	session := sessions.Default(c)
	session.Set("username", "张三111")
	//设置过期时间
	session.Options(sessions.Options{
		MaxAge: 30,
	})
	session.Save() //设置session时必须调用

	fmt.Println(models.UnixTotime(1629788418))
	//设置cookie,6个参数，名称，值，过期时间，生效路径、域，secure=true表示http中无效。但是https可以
	//Cookie保存在客户端浏览器，session保存在服务器
	c.SetCookie("username", "张三", 3600, "/", "localhost", false, true) //多个二级域名共享的话修改localhost就行

	//设置延时过期
	c.SetCookie("hobby", "吃饭 睡觉 打游戏", 15, "/", "localhost", false, true)
	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg": "DefaultController---我是一个msg",
		"t":   1629788418,
	})
}

func (def DefaultController) News1(c *gin.Context) {
	// username, _ := c.Cookie("username")
	// hobby, _ := c.Cookie("hobby")
	// c.String(200, "username=%v----hobby=%v\n", username, hobby)

	// c.String(200, "DefaultController---news1新闻\n")

	//session测试
	session := sessions.Default(c)
	username := session.Get("username")
	c.String(200, "username=%v\n", username)
}

func (def DefaultController) Shop(c *gin.Context) {
	username, _ := c.Cookie("username")
	hobby, _ := c.Cookie("hobby")
	c.String(200, "username=%v----hobby=%v\n", username, hobby)
	c.String(200, "DefaultController---Shop\n")

}

func (def DefaultController) DelerteCookie(c *gin.Context) {
	//删除cookie

	c.SetCookie("username", "张三", -1, "/", "localhost", false, true)
	c.String(200, "删除Cookie成功\n")
}
