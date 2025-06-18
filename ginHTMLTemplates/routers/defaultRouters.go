package routers

import (
	"ginframe/ginHTMLTemplates/controllers/yzgl"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", yzgl.DefaultController{}.Index)
		defaultRouters.GET("/news1", yzgl.DefaultController{}.News1)
		defaultRouters.GET("/shop", yzgl.DefaultController{}.Shop)
		defaultRouters.GET("/deletecookie", yzgl.DefaultController{}.DelerteCookie)
	}
}
