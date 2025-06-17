package main

//包括以下内容
//1.如何加载以及渲染模板
//2.多个模板文件夹，重名怎么办
//3.模板的常用语法
//3.1{{.}}来输出数据
//3.2 条件判断
//3.3 range循环遍历
//3.4 with解构结构体
//3.5 一些预定义函数 and or not len index 的使用,
//但是用法奇怪并且很多，一般自己封装自定义模板函数，就是在模板中封装一个函数
//自定义模板函数需要放在加载模板上面，配置路由下面
//4.模板的嵌套
//5.配置静态文件服务，引用了静态文件时，需要配置静态web服务r.Static("/static","./static")
//前面的/static表示路由，后面的表示路径
import (
	"fmt"
	"ginframe/ginHTMLTemplates/routers"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title string

	Content string
}

// 自定义模板函数
func UnixTotime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-2 15:04:05")
}
func Println(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + "---" + str2
}
func main() {
	r := gin.Default() //这个会默认有两个中间件
	//template.FuncMap是一个map[string]interface{}，字符串是方法名，后面的是方法实现
	r.SetFuncMap(template.FuncMap{
		"UnixTotime": UnixTotime,
		"Println":    Println,
	})
	//如果html模板有多个文件夹多个模块，这里写法要改变成templates/**/**
	//两个**表示一层目录
	r.LoadHTMLGlob("templates/**/*")
	//第一个参数表示路由，第二个表示映射的目录
	r.Static("/static", "./static") //匹配到XXXX路由时就会访问./static目录里的文件

	//路由分组抽离
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
			"score": 100,
			"hobby": []string{"吃饭", "睡觉", "写代码"},
			"newsList": []interface{}{
				&Article{
					Title:   "新闻标题111",
					Content: "111"},
				&Article{
					Title:   "新闻标题222",
					Content: "222"},
			},
			"testslice": []string{},
			"news": &Article{
				Title:   "新闻标题",
				Content: "新闻内容",
			},
			"date": 1749719852,
		})
	})
	//这时候admin和default都有index.html，不知道渲染哪一个
	//所以写模板的时候需要定义名称使用
	//{{define "admin/index.html"}} {{end}}
	r.GET("/news", func(c *gin.Context) {
		News := &Article{
			Title: "新闻标题",

			Content: "This is a news test",
		}
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻页面",
			"news":  News,
		})
	})
	//后台
	r.GET("/admin/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})

	r.GET("/admin/news", func(c *gin.Context) {

		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "后台新闻页面",
		})
	})

	r.Run()
}
