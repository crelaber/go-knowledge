package array

//题目：给定两个大小为m和n的有序数组nums1和nums2
//找出这两个有序数组的中位数，要求算法时间复杂度为0(log(m+n))

func FindMediaSortArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return FindMediaSortArrays(nums2, nums1)
	}

	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		nums1Mid = low + (high-low)>>1 //分界线右侧是mid，分界线左侧是mid+1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { //nums1中的分界线话多了，要向左移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] {
			low = nums1Mid + 1
		} else {
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	}
}
