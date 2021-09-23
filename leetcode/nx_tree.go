package leetcode

//N叉树的遍历

type NxTreeNode struct {
	Val      int
	Children []*NxTreeNode
}

//非递归法
func preOrder(root *NxTreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := []*NxTreeNode{root}
	for len(stack) > 0 {
		r := stack[len(stack)-1]
		stack := stack[:len(stack)-1]
		res = append(res, r.Val)
		var tmp []*NxTreeNode
		for _, v := range r.Children {
			tmp = append([]*NxTreeNode{v}, tmp...)
		}
		stack = append(stack, tmp...)
	}
	return res
}

//递归法
func preOrder2(root *NxTreeNode) []int {
	var res []int
	preOrderDfs(root, &res)
	return res
}

func preOrderDfs(root *NxTreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		for i := 0; i < len(root.Children); i++ {
			preOrderDfs(root.Children[i], res)
		}
	}
}
