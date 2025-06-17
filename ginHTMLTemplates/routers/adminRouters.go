package routers

import (
	"ginframe/ginHTMLTemplates/controllers/admin"
	"ginframe/ginHTMLTemplates/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	//只有注册了该中间件的分组路由或者路由才能get到set设置的数据
	adminRouters := r.Group("/admin", middlewares.InitMiddleware)
	//这个中间件只会作用于这个分组的路由
	//路由分组中配置中间件
	//第一种直接在adminRouters := r.Group("/admin"，XXX)路由后面添加
	//第二种，使用adminRouters.Use
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
