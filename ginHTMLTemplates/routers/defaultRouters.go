package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			//一样的方式渲染模板
			c.HTML(http.StatusOK, "default/index1.html", gin.H{
				"msg": "我是一个msg",
			})
		})
		defaultRouters.GET("/news1", func(c *gin.Context) {
			c.String(200, "news1新闻")
		})
	}
}
