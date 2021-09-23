package algo

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := []int{1, 9, 10, 30, 2, 5, 25, 29, 41, 234, 143, 12}
	max := SelectMax(arr)
	selectSort := SelectSort(arr)
	fmt.Println(max)
	fmt.Println(selectSort)
}
