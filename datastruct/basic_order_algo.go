package datastruct

//基础排序算法
//冒号排序
//现在有一堆乱序的数，比如：5 9 1 6 8 14 6 49 25 4 6 3。
//第一轮迭代：从第一个数开始，依次比较相邻的两个数，如果前面一个数比后面一个数大，那么交换位置，直到处理到最后一个数，最后的这个数是最大的。
//第二轮迭代：因为最后一个数已经是最大了，现在重复第一轮迭代的操作，但是只处理到倒数第二个数。
//第三轮迭代：因为最后一个数已经是最大了，最后第二个数是次大的，现在重复第一轮迭代的操作，但是只处理到倒数第三个数。
//第N轮迭代：….
//经过交换，最后的结果为：1 3 4 5 6 6 6 8 9 14 25 49，我们可以看到已经排好序了。
//因为小的元素会慢慢地浮到顶端，很像碳酸饮料的汽泡，会冒上去，所以这就是冒泡排序取名的来源。
func BubbleSort(list []int) {
	n := len(list)
	// 在一轮中有没有交换过
	didSwap := false
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			// 如果前面的数比后面的大，那么交换
			if list[j] > list[i] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}
		if !didSwap {
			return
		}
	}
}

//选择排序
//选择排序，一般我们指的是简单选择排序，也可以叫直接选择排序，它不像冒泡排序一样相邻地交换元素，而是通过选择最小的元素，每轮迭代只需交换一次。虽然交换次数比冒泡少很多，但效率和冒泡排序一样的糟糕。
//现在有一堆乱序的数，比如：5 9 1 6 8 14 6 49 25 4 6 3。
//第一轮迭代，从第一个数开始，左边到右边进行扫描，找到最小的数 1，与数列里的第一个数交换位置。
//第二轮迭代，从第二个数开始，左边到右边进行扫描，找到第二小的数 3，与数列里的第二个数交换位置。
//第三轮迭代，从第三个数开始，左边到右边进行扫描，找到第三小的数 4，与数列里的第三个数交换位置。
//第N轮迭代：….
//经过交换，最后的结果为：1 3 4 5 6 6 6 8 9 14 25 49，我们可以看到已经排好序了。
//每次扫描数列找出最小的数，然后与第一个数交换，然后排除第一个数，从第二个数开始重复这个操作，这种排序叫做简单选择排序
func SelectSort(list []int) {
	n := len(list)
	for i := 0; i < n-1; i++ {
		min := list[i]
		minIndex := i
		for j := i + 1; j < n; j++ {
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}
		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}

func SelectSortOptimize(list []int) {
	n := len(list)
	for i := 0; i < n/2; i++ {
		minIndex := i //最小值下标
		maxIndex := i //最大值下标
		// 在这一轮迭代中要找到最大值和最小值的下标
		for j := i + 1; j < n-i; j++ {
			if list[j] > list[maxIndex] {
				minIndex = j
				continue
			}
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}

		if maxIndex == i && minIndex != n-i-1 {
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}
