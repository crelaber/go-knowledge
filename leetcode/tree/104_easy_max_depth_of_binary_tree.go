package tree

import "go-knowledge/leetcode/utils"

//题目要求：输出一棵树的最大高度
//解题思路：遍历根节点的左孩子和右孩子的高度，取出两者的最大值加1即可
//时间复杂度o(n)，空间复杂度o(1)

func maxDepth(root *CommonTreeNode) int {
	if root == nil {
		return 0
	}
	return utils.Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
