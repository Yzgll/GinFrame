package api

import "github.com/gin-gonic/gin"

type ApiController struct {
}

func (api ApiController) Index(c *gin.Context) {
	c.String(200, "ApiController---我是一个api接口首页---")
}

func (api ApiController) UserList(c *gin.Context) {
	c.String(200, "ApiController---我是一个api接口-userlist---")
}

func (api ApiController) PList(c *gin.Context) {
	c.String(200, "ApiController---我是一个api接口-plist---")
}
