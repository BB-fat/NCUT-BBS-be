package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ncutbbs/module/oss"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func Upload(c *gin.Context) {
	// 单文件
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	filename := filepath.Base(file.Filename)
	filePath := fmt.Sprintf("./tmp/%s", filename)
	// 上传文件至指定目录
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	}
	fileUrl := oss.UploadImage(fmt.Sprintf("%d_%s", time.Now().Unix(), filename), filePath)
	c.String(http.StatusOK, fileUrl)
	_ = os.Remove(filePath)
}
