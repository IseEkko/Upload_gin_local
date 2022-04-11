package main

import (
	"UploadFileGin/Middler"
	"UploadFileGin/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.Use(Middler.Cors())
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.LoadHTMLGlob("templates/*")
	router.POST("/upload", controller.Upload)
	router.GET("/upload", controller.IndexHtmlUpload)
	router.POST("/QiNiuUpload", controller.QinuiYunUpLoad)
	router.Run(":8088")
}
