package qiniu

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

var (
	AccessKey = ""
	SecretKey = ""
	Bucket    = ""
	ImageUrl  = ""
)

func UploadFile(file multipart.File, fileSize int64) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: Bucket}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false, //收费
		UseHTTPS:      false, //收费

	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", "", err
	}
	fileName := ret.Key
	return ImageUrl, fileName, nil
}
