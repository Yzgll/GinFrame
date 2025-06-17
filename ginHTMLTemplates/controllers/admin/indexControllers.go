package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (Ic IndexController) Index(c *gin.Context) {
	username, _ := c.Get("username")
	fmt.Println(username)
	v, _ := username.(string)
	c.String(200, "IndexController--后台首页---"+v)
}
