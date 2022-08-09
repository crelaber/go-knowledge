package stack

//题目：下一个更大元素i 简单
//给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
//请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
//nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。

//示例 1:
//输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
//输出: [-1,3,-1]
//解释:
//对于 num1 中的数字 4 ，你无法在第二个数组中找到下一个更大的数字，因此输出 -1 。
//对于 num1 中的数字 1 ，第二个数组中数字1右边的下一个较大数字是 3 。
//对于 num1 中的数字 2 ，第二个数组中没有下一个更大的数字，因此输出 -1 。

//示例 2:
//输入: nums1 = [2,4], nums2 = [1,2,3,4].
//输出: [3,-1]
//解释:
//对于 num1 中的数字 2 ，第二个数组中的下一个较大数字是 3 。
//对于 num1 中的数字 4 ，第二个数组中没有下一个更大的数字，因此输出 -1 。

func bestNextGreaterElement(findNums []int, nums []int) []int {
	m := make(map[int]int)
	var s Stack
	for _, num := range nums {
		for !s.IsEmpty() && s.Peek() <= num { //将 nums 中的数依次进栈，一直向下找比栈顶值还大的数
			m[s.Pop()] = num
		}
		s.Push(num)
	}

	for i, num := range findNums {
		if greater, ok := m[num]; ok {
			findNums[i] = greater
		} else {
			findNums[i] = -1
		}
	}
	return findNums
}
