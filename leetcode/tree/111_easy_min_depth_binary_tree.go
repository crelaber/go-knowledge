package tree

import "go-knowledge/leetcode/utils"

//题目：给出一颗二叉树，找出其最小深度。最小深度是从根节点到最近叶子节点的最短路径上的节点数量，说明：叶子节点是没有子节点的节点
//解题思路：递归求出根节点到叶子节点的深度，输出最小深度即可
//复杂度：时间复杂度o(n)，空间复杂度o(1

func minDepth(root *CommonTreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return utils.Min(minDepth(root.Left), minDepth(root.Right)) + 1
}
