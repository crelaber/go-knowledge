package array

//题目大意：在数组中2个数之和等于给定值的数字，结果返回2个数字在数组的下标

//解题思路
//最优时间复杂度算法是o(n)
//顺序扫描数组，对每个元素，在map中招能组合给定值另一半数字，如果找到了，直接返回2个数字下标，如果找不到，就把这个数字存入map中，等待扫到另一半，在取出返回结果

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}
