package admin

import "github.com/gin-gonic/gin"

type ArticleController struct {
	BaseControllers
}

func (article ArticleController) Index(c *gin.Context) {
	c.String(200, "ArticleController---新闻列表首页---")
	article.Defeat(c)
}

func (article ArticleController) Add(c *gin.Context) {
	c.String(200, "ArticleController---增加新闻列表---")
}

func (article ArticleController) Edit(c *gin.Context) {
	c.String(200, "ArticleController---修改新闻列表---")
}
