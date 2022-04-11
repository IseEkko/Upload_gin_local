package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"time"
)

func Upload(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	newFileName := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000) + path.Ext(file.Filename)
	dst := "./file/" + newFileName
	// 上传文件至指定的完整文件路径
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		fmt.Println("err is :", err)
	}
	c.File(newFileName)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func IndexHtmlUpload(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "HtmlUploadIndex.html", gin.H{
		"title": "图片上传",
	})
}
