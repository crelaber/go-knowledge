package tree

import "strconv"

//题目：给定一个二叉树，返回所有从根节点到叶子节点的路径
//解题思路：递归求解
//复杂度：时间复杂度o(n)，空间复杂度o(1

func binaryTreePath(root *CommonTreeNode) []string {
	if root == nil {
		return []string{}
	}
	res := []string{}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	tmpLeft := binaryTreePath(root.Left)
	for i := 0; i < len(tmpLeft); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpLeft[i])
	}

	tmpRight := binaryTreePath(root.Right)
	for i := 0; i < len(tmpRight); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpRight[i])
	}
	return res
}
