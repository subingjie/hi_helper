package sort

// 选择排序
func SelectSort(data []int) {

	// 此层循环是为了找出data中最小的数
	for i := 0; i < len(data)-1; i++ {
		min := i
		for j := i + 1; j < len(data)-1; j++ {
			if data[i] > data[j] {
				min = j
			}
		}

		// 将两数交换
		data[i], data[min] = data[min], data[i]
	}

}
