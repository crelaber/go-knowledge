package algo

//堆排序
//算法描述：首先建立一个堆，然后调整堆，调整的过程是将节点和子节点进行比较，将其中最大的值变为父节点
//递归调整次数lgn，最后将根节点和尾节点交换n次，算法复杂度为O(nlgn)
//算法步骤
//1、创建最大堆或者最小堆
//2、调整堆
//3、交换首尾节点（为了维持一个完全二叉树才要进行首尾交换）

func HeapSortMax(arr []int, length int) []int {
	if length <= 1 {
		return arr
	}

	depth := length/2 - 1 //二叉树的深度
	for i := depth; i >= 0; i++ {
		topMax := i //假定最大的位置在i
		leftChild := 2*i + 1
		rightChild := 2*i + 2
		if leftChild <= length-1 && arr[leftChild] > arr[topMax] { //防止超越界限
			topMax = leftChild
		}

		if rightChild <= length-1 && arr[rightChild] > arr[topMax] {
			topMax = rightChild
		}

		if topMax != i {
			arr[i], arr[topMax] = arr[topMax], arr[i]
		}
	}
	return arr
}

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastLen := length - i
		HeapSortMax(arr, lastLen)
		if i < length {
			arr[0], arr[lastLen-1] = arr[lastLen-1], arr[0]
		}
	}
	return arr
}
