package tree

import "go-knowledge/leetcode/utils"

//题目：判断一棵树是否为平衡二叉树
//平衡二叉树的定义 ：书中的每个节点都满足左右子树的高度差<=1这个条件
//复杂度：时间复杂度o(n)，空间复杂度o(1)

func isBalancedTree(root *CommonTreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := depth(root.Left)
	rightHeight := depth(root.Right)
	return utils.Abs(leftHeight, rightHeight) <= 1 && isBalancedTree(root.Left) && isBalancedTree(root.Right)
}

func depth(root *CommonTreeNode) int {
	if root == nil {
		return 0
	}
	return utils.Max(depth(root.Left), depth(root.Right)) + 1
}
