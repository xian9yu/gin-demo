package ctrls

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

//单张图片上传
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	basePath := "D:\\Apps\\"
	filename := basePath + filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "文件上传成功 ",
	})
}

func Download(c *gin.Context) {
	filename, err := c.GetQuery("filename")
	if !err {
		c.String(200, "Success")
	}

	path := "./"
	path += filename
	fmt.Println(path)
	c.File(path)
}
