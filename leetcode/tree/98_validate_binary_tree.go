package tree

import "math"

//题目：给定一颗二叉树判断是否为有效的二叉搜索树
//二叉搜索树有如下特征
//1、节点的左子树包含小于当前节点的数
//2、节点的右子树包含大于当前节点的数
//3、所有左子树和右子树必须是二叉搜索树

func isValidBST(root *CommonTreeNode) bool {
	return isValidBst1(root, math.Inf(-1), math.Inf(1))
}

func isValidBst1(root *CommonTreeNode, min, max float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)
	return v < max && v > min && isValidBst1(root.Left, min, v) && isValidBst1(root.Right, v, max)
}

//解法2，把bst左中右的顺序输出到数组中，如果是BST，则数组是有序的，如果逆序就不是
func isValidBSt2(root *CommonTreeNode) bool {
	arr := []int{}
	treeInOrder(root, &arr)
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func treeInOrder(root *CommonTreeNode, arr *[]int) {
	if root == nil {
		return
	}
	treeInOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	treeInOrder(root.Right, arr)
}
