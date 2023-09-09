package tencent

import (
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
func Upload(client *cos.Client, localFile string, cosFile string) error {

	_, err := client.Object.PutFromFile(nil, cosFile, localFile, nil)

	return err
}

func GetCosFile() {

}
