package admin

import "github.com/gin-gonic/gin"

type IndexController struct {
}

func (Ic IndexController) Index(c *gin.Context) {
	c.String(200, "IndexController--后台首页---")
}
