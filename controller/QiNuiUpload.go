package controller

import (
	"UploadFileGin/errmsg"
	"UploadFileGin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"net/http"
)

func QinuiYunUpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	NewQiNui := model.NewQiNui()
	files, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误,请检查文件路径:", err)
	}
	NewQiNui.LoadQiniu(files)
	url, code := NewQiNui.UpLoadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.SUCCESS,
		"url":     url,
	})
}
