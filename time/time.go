package hitime

import "time"

func ParseStrTimeToInt(timeStr string) (int64, error) {
	// 将timeStr转换成时间对象
	timeObj, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return 0, err
	}
	return timeObj.Unix(), nil
}
