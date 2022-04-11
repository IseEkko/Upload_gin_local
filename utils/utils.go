package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func Init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误,请检查文件路径:", err)
	}
	LoadQiniu(file)
}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()
	fmt.Println(AccessKey, SecretKey, Bucket)
}
