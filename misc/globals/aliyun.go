package globals

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"fmt"
)

const (
	endpoint = "https://oss-cn-hangzhou.aliyuncs.com"
	accessKeyId     = "LTAIovfYZb1LbExG"
	accessKeySecret = "RKHGrsiDA6bu96RqdPLBq1DG0ZjvAR"
	bucket = "siiva-video"
	bucket_public = "siiva-video-public"

)

func PutFile(localPath string, remotePath string) error {

	client, err1 := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err1 != nil {
		fmt.Println("err1:", err1)
		return err1
	}

	bucket, err2 := client.Bucket(bucket)
	if err2 != nil {
		fmt.Println("err2:", err2)
		return err2
	}

	err3 := bucket.PutObjectFromFile(remotePath, localPath)
	if err3 != nil {
		fmt.Println("err3:", err3)
		return err3
	}

	fmt.Println(localPath + "上传成功")
	return nil
}

func PutFilePublic(localPath string, remotePath string) error {

	client, err1 := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err1 != nil {
		fmt.Println("err1:", err1)
		return err1
	}

	bucket, err2 := client.Bucket(bucket_public)
	if err2 != nil {
		fmt.Println("err2:", err2)
		return err2
	}

	err3 := bucket.PutObjectFromFile(remotePath, localPath)
	if err3 != nil {
		fmt.Println("err3:", err3)
		return err3
	}

	fmt.Println(localPath + "上传成功")
	return nil
}