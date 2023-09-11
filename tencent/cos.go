package tencent

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

// NewClient 创建cos客户端
func NewClient(buckName string, appId string, secretId string, secretKey string, region string) (*cos.Client, error) {

	cosUrlFormat := "http://%s-%s.cos.%s.myqcloud.com"

	bucketUrl, urlParseErr := url.Parse(fmt.Sprintf(cosUrlFormat, buckName, appId, region))

	if urlParseErr != nil {
		return nil, urlParseErr
	}

	baseUrl := &cos.BaseURL{
		BucketURL: bucketUrl,
	}

	client := cos.NewClient(baseUrl, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretId,
			SecretKey: secretKey,
		},
	})

	return client, nil
}

// Upload 上传文件
func Upload(data []byte, cosPath string, client *cos.Client) error {

	file := bytes.NewReader(data)

	_, err := client.Object.Put(context.Background(), cosPath, file, nil)
	if err != nil {
		return err
	}
	return nil
}

// GetCosFile 获取cos文件数据
func GetCosFile() ([]byte, error) {
	return nil, nil
}
