package admin

import "github.com/gin-gonic/gin"

type BaseControllers struct {
}

func (base BaseControllers) Sucess(c *gin.Context) {
	c.String(200, "成功，BaseControllers用来演示继承")
}

func (base BaseControllers) Defeat(c *gin.Context) {
	c.String(200, "失败，BaseControllers用来演示继承")
}
