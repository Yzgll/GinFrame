package admin

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseControllers
}

// 采用结构体可以继承
func (con UserController) Index(c *gin.Context) {
	c.String(200, "UserController--用户列表---\n")
	//c.String(200, "这是用户列表页面，UserController继承了BaseControllers，调用了方法\n")
	// con.Sucess(c)

}

func (con UserController) Add(c *gin.Context) {
	//c.String(200, "UserController--添加用户---\n")
	c.HTML(http.StatusOK, "admin/useradd.html", gin.H{}) //上传单个文件
}
func (con UserController) Add1(c *gin.Context) {
	//c.String(200, "UserController--添加用户---\n")
	c.HTML(http.StatusOK, "admin/useradd1.html", gin.H{}) //上传多个重名文件
}
func (con UserController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useredit.html", gin.H{}) //上传多个文件
}

func (con UserController) DoUpLoad(c *gin.Context) {
	//执行真正的上传
	//获取表单的username
	username := c.PostForm("username")
	//获取文件
	file, err := c.FormFile("face")
	//获取文件名file.Filename  假如是aaa.jpg
	dst := path.Join("./static/upload", file.Filename) //拼接保存路径
	if err == nil {
		//dst是文件要保存的目录，一般是静态

		//拼接好就是 ./static/upload/aaa.jpg
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"sucess":   true,
		"username": username,
		"dst":      dst,
	})
}

func (con UserController) DoUpLoad1(c *gin.Context) {
	//执行真正的上传
	//获取表单的username
	username := c.PostForm("username")
	//获取表单
	form, _ := c.MultipartForm()
	files := form.File["face[]"]

	//遍历获取的多个文件
	for _, file := range files {
		//上传
		dst := path.Join("./static/upload", file.Filename)
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"sucess":   true,
		"username": username,
	})
}
func (con UserController) DoUpLoads(c *gin.Context) {
	//执行真正的上传
	//获取表单的username
	username := c.PostForm("username")
	//获取多个文件
	file1, err1 := c.FormFile("face1")
	file2, err2 := c.FormFile("face2")
	//获取文件名file.Filename  假如是aaa.jpg
	dst1 := path.Join("./static/upload", file1.Filename) //拼接保存路径
	dst2 := path.Join("./static/upload", file2.Filename)
	if err1 == nil {
		//dst是文件要保存的目录，一般是静态

		//拼接好就是 ./static/upload/aaa.jpg
		c.SaveUploadedFile(file1, dst1)

	}

	if err2 == nil {
		//dst是文件要保存的目录，一般是静态

		//拼接好就是 ./static/upload/aaa.jpg
		c.SaveUploadedFile(file2, dst2)

	}
	c.JSON(http.StatusOK, gin.H{
		"sucess":   true,
		"username": username,
		"dst1":     dst1,
		"dst2":     dst2,
	})

}
