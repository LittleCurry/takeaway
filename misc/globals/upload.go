package globals

import (
	"io"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
	ALIYUNOSSKEY     = "LTAIAHvYCJg3q0sp"
	ALIYUNOSSSECRET = "EOPCMRjjW3mDC8MFV4LSwAMEiMKVny"
)

func UploadImg(imgName string, img io.Reader ) (string, error) {

	bucketName := "siiva-video-public"
	client, err1 := oss.New("http://oss-cn-hangzhou.aliyuncs.com", ALIYUNOSSKEY, ALIYUNOSSSECRET)
	if err1 != nil {
		fmt.Println("err1:", err1)
		return "", err1
	}
	bucket, err2 := client.Bucket(bucketName)
	if err2 != nil {
		fmt.Println("err2:", err2)
		return "", err2
	}
	// 上传文件
	err3 := bucket.PutObject(imgName, img)
	if err3 != nil {
		fmt.Println("err3:", err3)
	}
	// 保存到磁盘
	//src, err := file.Open()
	//defer src.Close()
	//out, err := os.Create("./uploads/"+file.Filename)
	//defer out.Close()
	//io.Copy(out, src)
	//err3 := c.SaveUploadedFile(file, "./uploads/"+file.Filename)
	return "https://siiva-video-public.oss-cn-hangzhou.aliyuncs.com/"+imgName, nil
}
