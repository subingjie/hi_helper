package histring

import (
	"math"
	"strconv"
	"time"
)

// strToTime 字符串转时间
func strToTime(str string) (time.Time, error) {
	// 2006-01-02 15:04:05
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, str)

	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// strToTime 字符串转float64，保留num位小数
func strToFload64(str string, num int) (float64, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}

	// 将f保留num位小数
	f = float64(int(f*math.Pow10(num))) / math.Pow10(num)
	return f, nil
}

// strToInt64 字符串转int64
// func strToInt64(str string) (int64, error) {
// 	i, err := strconv.ParseInt(str, 10, 64)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return i, nil
// }
