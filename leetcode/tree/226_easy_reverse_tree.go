package tree

//题目：反转二叉树
//难度：easy
//解题思路：使用递归解决，先递归调用反转根节点的左孩子，然后递归反转根节点的右孩子，然后左右交换根节点的左孩子和右孩子

func invertTree(root *CommonTreeNode) *CommonTreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
