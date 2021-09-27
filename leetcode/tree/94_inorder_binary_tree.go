package tree

//题目：中序遍历二叉树

func BinaryTreeInOrderTraverse(root *CommonTreeNode) []int {
	var result []int
	inOrder(root, &result)
	return result
}

func inOrder(root *CommonTreeNode, output *[]int) {
	if root != nil {
		inOrder(root.Left, output)
		*output = append(*output, root.Val)
		inOrder(root.Right, output)
	}
}
