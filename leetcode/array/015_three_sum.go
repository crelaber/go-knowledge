package array

import "sort"

//题目：给定一个数组，要求这个数组中找出3个数之和为0的所有组合
//解题思路：用map提前计算好任意2个数字之和，保存起来，可以将时间复杂度江伟O(n^2)，比较麻烦的点在于，最后输出解的时候，要求输出不重复的解
//数组中同一个数字可能出现多次，同一个数字可能使用多次，不能重复。例如【-1，-1，2】、【2，-1，-1】、【-1，2，-1】这3个解是重读的

//这里需要去重和排序，map记录每个数字出现的次数，然后对map的key数组 进行排序，最后这个排序以后的数组里面扫描，找到另外两个数字和自己组成0的组合

func ThreeNum(nums []int) [][]int {
	sort.Ints(nums)
	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	for index = 1; index < length-1; index++ {
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}

			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}
