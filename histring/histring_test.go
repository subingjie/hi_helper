package histring

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {

}

func romanToInt(s string) int {

	var strMap = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var num int

	for i, char := range s {
		value := strMap[fmt.Sprint(char)]
		if i+1 < len(s) {
			num -= value
		} else {
			num += value
		}
	}

	return num
}
