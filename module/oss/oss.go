package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"ncutbbs/config"
	"os"
)

var client *oss.Client
var bucket *oss.Bucket

func InitOSS() {
	var err error
	client, err = oss.New(config.Config.OSSEndpoint, config.Config.OSSAccessKeyID, config.Config.OSSAccessKeySecret)
	if err != nil {
		fmt.Println("OSS 初始化失败:", err)
		os.Exit(-1)
	}

	bucket, _ = client.Bucket("ncut-bbs")
}

func UploadImage(filename, filePath string) string {
	err := bucket.PutObjectFromFile(filename, filePath)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("%s/%s", config.Config.OSSBaseURL, filename)
}
