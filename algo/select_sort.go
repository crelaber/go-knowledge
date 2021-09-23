package algo

//选择排序
//算法描述：从未排序数据中选择最大或者最小值和当前值交换，算法复杂度O(n^2)

//算法步骤：
//1、选择一个数当最小值或者最大值，进行比较然后交换
//2、循环向后查询

////获取切片里面的最大值
func SelectMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	}
	max := arr[0]
	for i := 0; i < length; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func SelectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 1; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}
