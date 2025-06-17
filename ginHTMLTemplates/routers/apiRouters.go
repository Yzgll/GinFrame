package routers

import (
	"ginframe/ginHTMLTemplates/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {

	apiRouters := r.Group("/api")
	{

		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.GET("/userlist", api.ApiController{}.UserList)
		apiRouters.GET("/plist", api.ApiController{}.PList)
	}
}
