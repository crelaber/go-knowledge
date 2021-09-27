package tree

//题目：给定一个整数，生成所有1...n为节点组成的二叉搜索树
//解题思路：输出1-n组成的BST的解

func generateTrees(n int) []*CommonTreeNode {
	if n == 0 {
		return []*CommonTreeNode{}
	}
	return generateBsTree(1, n)
}

func generateBsTree(start, end int) []*CommonTreeNode {
	tree := []*CommonTreeNode{}
	if start > end {
		tree = append(tree, nil)
		return tree
	}
	for i := start; i <= end; i++ {
		left := generateBsTree(start, i-1)
		right := generateBsTree(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root := &CommonTreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				}
				tree = append(tree, root)
			}
		}
	}
	return tree
}
