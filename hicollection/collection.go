package hicollection

import (
	"encoding/json"
	"sort"
	"time"
)

/*
集合处理工具方法包
*/

// ListUniqueForStr 字符串列表去重
func ListUniqueForStr(elements []string) []string {
	// 创建一个map用来存放不重复的元素
	seen := make(map[string]bool)

	// result用来存放返回的不重复切片
	result := []string{}

	// 循环切片
	for _, str := range elements {
		if !seen[str] {
			seen[str] = true
			result = append(result, str)
		}
	}

	return result
}

// ListUniqueForInterface 任意类型列表去重
func ListUniqueForInterface(elements []interface{}) []interface{} {
	seen := make(map[interface{}]struct{})
	result := []interface{}{}

	for _, elem := range elements {
		if _, ok := seen[elem]; !ok {
			seen[elem] = struct{}{}
			result = append(result, elem)
		}
	}

	return result
}

// PraseDataToStruct 将数据转换成结构体
func PraseDataToStruct(originData interface{}, targetStruct interface{}) error {
	// 首先判断originData是不是[]byte类型
	jsonOriginData, ok := originData.([]byte)

	// 如果不是[]byte类型，就转换成[]byte类型
	if !ok {
		var jsonErr error
		jsonOriginData, jsonErr = json.Marshal(originData)
		if jsonErr != nil {
			return jsonErr
		}
	}

	// 进行json处理
	if err := json.Unmarshal(jsonOriginData, targetStruct); err != nil {
		return err
	}

	return nil
}

// sortListByInt 根据int类型的key对map进行排序
func sortListByIntKey(data []map[string]interface{}, key string) []map[string]interface{} {

	sort.SliceStable(data, func(i, j int) bool {
		value1 := data[i][key].(int64)
		value2 := data[j][key].(int64)

		return value1 < value2
	})

	return data
}

// sortListByTimeStrKey 根据时间字符串类型的key对map进行排序
func sortListByTimeStrKey(data []map[string]interface{}, key string) []map[string]interface{} {
	sort.SliceStable(data, func(i, j int) bool {

		time1, _ := time.Parse("2006-01-02 15:04:05", data[i][key].(string))
		time2, _ := time.Parse("2006-01-02 15:04:05", data[j][key].(string))

		return time1.Before(time2)
	})

	return data
}
