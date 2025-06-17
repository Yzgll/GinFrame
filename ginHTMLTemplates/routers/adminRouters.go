package routers

import (
	"ginframe/ginHTMLTemplates/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", admin.IndexController{}.Index)

		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/admin/user/add", admin.UserController{}.Add) //调用控制其的方法，不加括号，只注册，不然表示执行
		adminRouters.GET("/admin/user/edit", admin.UserController{}.Edit)

		adminRouters.GET("/article", admin.ArticleController{}.Index)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/edit", admin.ArticleController{}.Edit)
	}
}
