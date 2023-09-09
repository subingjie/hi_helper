package tencent

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestTencent(t *testing.T) {

	c, err := NewClient("su-default", "1315783504", "AKIDLbgbJn4CAyZv53rTuCINGySksmX5QXPi", "CBykLkwysb9WilmiHtLfpmZ9NQPHEGoJ", "ap-guangzhou")

	fmt.Printf("err: %v\n", err)

	// fileData 为模拟文件的二进制数据
	fileData := []byte{}

	// 读取本地的test.text文件，并将其内容写入到fileData中
	fileData, _ = ioutil.ReadFile("test.txt")

	// 将fileData，写入到io.Reader中
	file := bytes.NewReader(fileData)

	c.Object.Put(context.Background(), "test.txt", file, nil)

}
