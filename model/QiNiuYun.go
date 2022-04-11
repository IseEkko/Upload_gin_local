package model

import (
	"UploadFileGin/errmsg"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"gopkg.in/ini.v1"
	"mime/multipart"
)

type QiNui struct {
	AccessKey string
	SecretKey string
	Bucket    string
	ImgUrl    string
}

func (this *QiNui) LoadQiniu(file *ini.File) {

	this.AccessKey = file.Section("").Key("AccessKey").String()
	this.SecretKey = file.Section("").Key("SecretKey").String()
	this.Bucket = file.Section("").Key("Bucket").String()
	this.ImgUrl = file.Section("").Key("QiniuServer").String()

}
func (this *QiNui) UpLoadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: this.Bucket,
	}
	mac := qbox.NewMac(this.AccessKey, this.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	fmt.Println(err)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := this.ImgUrl + ret.Key
	return url, errmsg.SUCCESS
}

func NewQiNui() *QiNui {
	return &QiNui{}
}
