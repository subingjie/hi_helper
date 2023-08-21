package hicollection

import (
	"fmt"
	"testing"
)

func TestCollection(t *testing.T) {
	testListUniqueForStr()
}

func testListUniqueForStr() {
	elements := []string{"1", "2", "3", "4", "4", "5", "5", "6"}
	result := ListUniqueForStr(elements)
	fmt.Printf("result: %v\n", result)
}
