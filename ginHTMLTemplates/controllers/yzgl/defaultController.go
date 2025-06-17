package yzgl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func (def DefaultController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "default/index1.html", gin.H{
		"msg": "DefaultController---我是一个msg",
	})
}

func (def DefaultController) News1(c *gin.Context) {
	c.String(200, "DefaultController---news1新闻")
}
