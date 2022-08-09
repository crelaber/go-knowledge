package tree

//题目：给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上的节点值相加等于目标和，说明叶子节点是没有子节点的节点
//解题思路：递归求解
//复杂度：时间复杂度o(n)，空间复杂度o(1

func hasPathSum(root *CommonTreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
