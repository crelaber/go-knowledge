package tree

// IsSameTree 判断 2 颗树是否是完全相等的。
//难度：easy
//解题思路：通过递归即可
//时间复杂度o(n)，空间复杂度o(1)
func IsSameTree(p *CommonTreeNode, q *CommonTreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
	}
	return false
}
