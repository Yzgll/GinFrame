package admin

import "github.com/gin-gonic/gin"

type UserController struct {
	BaseControllers
}

//采用结构体可以继承
func (con UserController) Index(c *gin.Context) {
	c.String(200, "UserController--用户列表---")
	c.String(200, "这是用户列表页面，UserController继承了BaseControllers，调用了方法")
	con.Sucess(c)

}

func (con UserController) Add(c *gin.Context) {
	c.String(200, "UserController--增加用户---")
}

func (con UserController) Edit(c *gin.Context) {
	c.String(200, "UserController--修改用户---")
}
