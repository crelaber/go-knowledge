package tree

//题目大意：判断二叉树是否为左右对称
//难度：easy
//解题思路：
//1.将根节点的左子树反转二叉树，然后再和根节点和有节点进行比较，是否完全相等
// 2. 反转二叉树是226题，判断二叉树是否相等为100题目
//时间复杂度o(n)，空间复杂度o(1)

//解法一dfs
func isSymmetric(root *CommonTreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left *CommonTreeNode, right *CommonTreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}
	return (left.Val == right.Val) && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

//解法二，先将左子树反转，然后再判断是否和右子树相等
func isSymmetric2(root *CommonTreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree2(invertTree2(root.Left), root.Right)
}

// 判断是否相等，第100题
func isSameTree2(p *CommonTreeNode, q *CommonTreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree2(p.Left, q.Left) && isSameTree2(p.Right, q.Right)
	}

	return false
}

//反转二叉树，第226题
func invertTree2(root *CommonTreeNode) *CommonTreeNode {
	if root == nil {
		return nil
	}
	invertTree2(root.Left)
	invertTree2(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
