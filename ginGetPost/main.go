package main

//get传值，从网址中获取数据
//获取get post传递的参数数据绑定到结构体
//动态路由传值
import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 配置多个tag
type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Article struct {
	Title   string `json:"title" form:"title" xml:"title"`
	Content string `json:"content" form:"content" xml:"content"`
}

func main() {
	r := gin.Default()
	//get请求传值
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1") //但你没有值时返回默认值1
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})

	r.GET("/article", func(c *gin.Context) {
		id := c.DefaultQuery("id", "1")
		c.JSON(http.StatusOK, gin.H{
			"msg": "新闻详情",
			"id":  id,
		})
	})

	//post传值
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/user.html", gin.H{})
	})
	//获取表单post的数据
	r.POST("/doAddUsr1", func(c *gin.Context) {
		username := c.PostForm("username") //需要和表单定义的name相同
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	r.GET("/getuser", func(c *gin.Context) {
		user := &UserInfo{}
		//http://localhost:8080/getuser?username=zhangsan&password=abc
		//将get或者post得到的数据解析到结构体中
		err := c.ShouldBind(&user)
		fmt.Printf("%#v", user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		} else {
			//成功返回user结构体
			c.JSON(http.StatusOK, user)
		}
	})
	//post的值也可以绑定到结构体
	r.POST("/doAddUsr2", func(c *gin.Context) {
		user := &UserInfo{}
		err := c.ShouldBind(&user)
		fmt.Printf("%#v", user)
		if err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})
	//获取post xml数据
	r.POST("/xml", func(c *gin.Context) {
		xmlSliceData, _ := c.GetRawData() //从c.Request.Body读取请求数据
		//获取到数据存放在xmlSliceData中将切片转化为结构体
		article := &Article{}
		err := xml.Unmarshal(xmlSliceData, &article)
		if err == nil {
			//成功了返回数据
			c.JSON(http.StatusOK, article)
			fmt.Printf("%#v", article)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}

	})

	//动态路由传值
	// localhost:8080/list/123,就会把123赋值给cid
	r.GET("/list/:cid", func(c *gin.Context) {
		//获取方法是
		cid := c.Param("cid")
		c.String(200, "%v", cid)
	})
	r.Run()
}
