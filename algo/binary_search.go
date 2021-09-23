package algo

import (
	"fmt"
)

//二分查找算法
//算法描述：在一组有序数组中，将数组一分为二，将要查询的元素和风格店进行比较，分为三种情况
//1、相等直接返回
//2、元素大于分隔点，在分隔点右侧继续查找
//3、元素小于分隔点，在分隔点左侧继续查找
//时间复杂度O(lgn)

//要求：必须是有序数组，并能支持随机访问

//变形：
//1、在查找第一个值等于给定的，在相等的时候做处理，向前查
//2、查找最后一个值等于给定的值，在相等的时候做处理，向后查
//3、查找第一个大于等于给定的值，判断边界减1
//4、查找最后一个小于等于给定的值，判断边界加1

//实际应用
//1、用户ip区间段查询
//2、用于相似度查询

func BinarySearch(arr []int, findData int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		fmt.Println(mid)
		if arr[mid] > findData {
			high = mid - 1
		} else if arr[mid] < findData {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
