package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Au    string `json:"au"`
	Price int    `json:"price"`
	Title string `json:"title"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//第一种返回字符串相应
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "这是第一种返回的是%v响应信息", "字符串") //三个参数，其实也可以只有两个不写第三个
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"这是map的键":      "这是map的值，是一个空接口",
			"ResponseType": "JSON",
		})
	})
	//JSON也可以返回结构体
	r.GET("/json1", func(c *gin.Context) {
		a := &Book{
			Au:    "Yzgl",
			Price: 100,
			Title: "这是一本名著",
		}
		c.JSON(http.StatusOK, a)
	})
	//通过tag标签可以返回指定的格式
	r.GET("/json2", func(c *gin.Context) {
		a := &Book{
			Au:    "Yzgl",
			Price: 100,
			Title: "这是一本名著",
		}
		c.JSON(http.StatusOK, a)
	})
	//相应JSONP请求，JSONP和JSON差不多，但是JSONP可以传入回调函数
	//localhost:8080/jsonp?callback=xxx 这就是回调函数
	//输出为xxx({ "au": "Yzgl-JSONP", "price": 100, "title": "这是一本名著" });，会把数据放在回调函数xxx中
	r.GET("/jsonp", func(c *gin.Context) {
		a := &Book{
			Au:    "Yzgl-JSONP",
			Price: 100,
			Title: "这是一本名著",
		}
		c.JSONP(http.StatusOK, a)
	})

	//相应一个xml数据
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"ResponseType": "XML",
			"sucess":       true,
		}) //可以传入map[string]interface{}，可以和json一样传入结构体
	})
	//需要配r.LoadHTMLGlob("templates/*")
	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台数据",
		})
	})

	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title":     "我是后台数据",
			"goodsType": "Phone",
			"price":     200,
		})
	})
	r.Run()
}
